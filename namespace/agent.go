package namespace

import (
	"errors"
	"fmt"
	m2 "github.com/behavioral-ai/collective/messaging"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	namespaceName   = "core:agent/collective/namespace"
	defaultDuration = time.Second * 10
)

var (
	agent *agentT
)

type agentT struct {
	running   bool
	duration  time.Duration
	relations *relationT
	using     map[string]m2.UsingRecord

	//handler  eventing.Agent
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

func NewAgent() messaging.Agent {
	agent = newAgent()
	return agent
}

func newAgent() *agentT {
	a := new(agentT)
	a.duration = defaultDuration
	a.using = make(map[string]m2.UsingRecord)
	a.relations = newRelation()

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent name
func (a *agentT) Name() string { return namespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if !a.running {
		if m.Name == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Name == messaging.StartupEvent {
			a.run()
			a.running = true
			return
		}
		return
	}
	if m.Name == messaging.ShutdownEvent {
		a.running = false
	}
	switch m.Channel() {
	case messaging.ChannelEmissary:
		a.emissary.Send(m)
	case messaging.ChannelMaster:
		a.master.Send(m)
	case messaging.ChannelControl:
		a.emissary.Send(m)
		a.master.Send(m)
	default:
		a.emissary.Send(m)
	}
}

func (a *agentT) configure(m *messaging.Message) {
	if use, ok := m2.UsingContent(m); ok {
		a.using[use.Collective] = use
	}
	//messaging.Reply(m, messaging.StatusOK(), a.Uri())
}

// Run - run the agent
func (a *agentT) run() {
	go masterAttend(a)
	go emissaryAttend(a)

}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}

func (a *agentT) addThing(name, cname, authority, author string) *messaging.Status {
	if name == "" || authority == "" {
		return messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v or authority %v", name, authority)))
	}
	_, status := httpPutThing(name, cname, authority, author)
	if !status.OK() {
		return status.WithMessage(fmt.Sprintf("name %v", name))
	}
	return status
}

func (a *agentT) addRelation(name, cname, thing1, thing2, author string) *messaging.Status {
	if name == "" || thing1 == "" || thing2 == "" {
		return messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name1 %v or name2 %v or author %v", thing1, thing2, author)))
	}
	_, status := httpPutRelation(name, cname, thing1, thing2, author)
	if !status.OK() {
		return status.WithMessage(fmt.Sprintf("name1 %v", name))
	}
	a.relations.put(name, thing1, thing2)
	return status
}
