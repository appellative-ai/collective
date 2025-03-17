package namespace

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	AgentNamespaceName = "resiliency:agent/behavioral-ai/collective/namespace"
	agentUri           = AgentNamespaceName
	defaultDuration    = time.Second * 10
)

type agentT struct {
	running  bool
	agentId  string
	hostName string
	uri      []string
	duration time.Duration

	ticker     *messaging.Ticker
	emissary   *messaging.Channel
	master     *messaging.Channel
	notifier   messaging.NotifyFunc
	dispatcher messaging.Dispatcher
}

func newAgent(dispatcher messaging.Dispatcher) *agentT {
	a := new(agentT)
	a.agentId = agentUri
	a.duration = defaultDuration
	a.ticker = messaging.NewTicker(messaging.Emissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	a.dispatcher = dispatcher
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return a.agentId }

// Name - agent name
func (a *agentT) Name() string { return AgentNamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.Emissary:
		a.emissary.Send(m)
	case messaging.Master:
		a.master.Send(m)
	case messaging.Control:
		a.emissary.Send(m)
		a.master.Send(m)
	default:
		a.emissary.Send(m)
	}
}

// Run - run the agent
func (a *agentT) Run() {
	if a.running {
		return
	}
	go masterAttend(a)
	go emissaryAttend(a)
	a.running = true
}

// Shutdown - shutdown the agent
func (a *agentT) Shutdown() {
	if !a.emissary.IsClosed() {
		a.emissary.Send(messaging.Shutdown)
	}
	if !a.master.IsClosed() {
		a.master.Send(messaging.Shutdown)
	}
}

func (a *agentT) notify(e messaging.Event) {
	if a.notifier != nil {
		a.notifier(e)
	} else {
		httpNotify(e)
	}
}

func (a *agentT) dispatch(channel any, event string) {
	messaging.Dispatch(a, a.dispatcher, channel, event)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}

func (a *agentT) addThing(nsName, author string) *messaging.Status {
	if nsName == "" || author == "" {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v or author %v", nsName, author)), a.Uri())
	}
	_, status := httpPutThing(nsName, author)
	if !status.OK() {
		status.SetAgent(a.Uri())
		status.SetMessage(fmt.Sprintf("name %v", nsName))
		return status
	}
	return status
}

func (a *agentT) addRelation(nsName1, nsName2, author string) *messaging.Status {
	if nsName1 == "" || author == "" || nsName2 == "" {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name1 %v or name2 %v or author %v", nsName1, nsName2, author)), a.Uri())
	}
	_, status := httpPutRelation(nsName1, nsName2, author)
	if !status.OK() {
		status.SetAgent(a.Uri())
		status.SetMessage(fmt.Sprintf("name1 %v", nsName1))
		return status
	}
	return status
}
