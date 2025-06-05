package namespace

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/collective/private"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	NamespaceAgentName = "common:core:agent/namespace/collective"
	defaultDuration    = time.Second * 10
)

var (
	agent *agentT
)

type agentT struct {
	running   bool
	duration  time.Duration
	relations *relationT
	intf      *private.Interface

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
	a.relations = newRelation()
	a.intf = private.NewInterface()

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent name
func (a *agentT) Name() string { return NamespaceAgentName }

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
	switch m.ContentType() {
	case private.ContentTypeInterface:
		intf, status := private.InterfaceContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
		}
		a.intf = intf
	}
	messaging.Reply(m, messaging.StatusOK(), a.Name())
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

func (a *agentT) addThing(name, cname, author string) *messaging.Status {
	if name == "" || author == "" {
		return messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v or authority %v", name, author)))
	}
	status := a.intf.Thing(http.MethodPut, name, cname, author)
	if !status.OK() {
		return status.WithMessage(fmt.Sprintf("name %v", name))
	}
	return status
}

func (a *agentT) addRelation(name, cname, thing1, thing2, author string) *messaging.Status {
	if name == "" || thing1 == "" || thing2 == "" || author == "" {
		return messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name1 %v or name2 %v or author %v", thing1, thing2, author)))
	}
	// TODO: remove after initial testing
	a.relations.put(name, thing1, thing2)

	status := a.intf.Relation(http.MethodPut, name, cname, thing1, thing2, author)
	if !status.OK() {
		return status.WithMessage(fmt.Sprintf("name1 %v", name))
	}
	return status
}
