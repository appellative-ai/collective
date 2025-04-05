package namespace

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	AgentNamespaceName = "resiliency:agent/behavioral-ai/collective/namespace"
	defaultDuration    = time.Second * 10
)

type agentT struct {
	running  bool
	duration time.Duration

	handler  messaging.Agent
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

func newAgent(handler messaging.Agent) *agentT {
	a := new(agentT)
	a.duration = defaultDuration
	if handler != nil {
		a.handler = handler
	} else {
		a.handler = eventing.Agent
	}
	a.ticker = messaging.NewTicker(messaging.Emissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return AgentNamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Event() == messaging.ConfigEvent {
		a.configure(m)
		return
	}
	if m.Event() == messaging.StartupEvent {
		a.run()
		return
	}
	if !a.running {
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

func (a *agentT) configure(m *messaging.Message) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
	}
	// configure
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}

// Run - run the agent
func (a *agentT) run() {
	if a.running {
		return
	}
	go masterAttend(a)
	go emissaryAttend(a)
	a.running = true
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
		status.WithAgent(a.Uri())
		status.WithMessage(fmt.Sprintf("name %v", nsName))
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
		status.WithAgent(a.Uri())
		status.WithMessage(fmt.Sprintf("name1 %v", nsName1))
		return status
	}
	return status
}
