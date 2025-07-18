package resource

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/collective/private"
	"github.com/appellative-ai/core/messaging"
	"net/http"
	"time"
)

const (
	NamespaceName   = "common:core:agent/resource/collective"
	defaultDuration = time.Second * 10
)

var (
	agent *agentT
)

type text struct {
	Value string
}

type agentT struct {
	running  bool
	duration time.Duration
	cache    *cacheT
	intf     *private.Interface

	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

func NewAgent() messaging.Agent {
	if agent == nil {
		agent = newAgent()
	}
	return agent
}

func newAgent() *agentT {
	a := new(agentT)
	a.duration = defaultDuration
	a.cache = newCache()
	a.intf = private.NewInterface()

	//a.handler = handler
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}

	switch m.Name {
	case messaging.ConfigEvent:
		if a.running {
			return
		}
		a.configure(m)
		return
	case messaging.StartupEvent:
		if a.running {
			return
		}
		a.run()
		a.running = true
		return
	case messaging.ShutdownEvent:
		if !a.running {
			return
		}
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

func (a *agentT) getRepresentation(name string) (messaging.Content, *messaging.Status) {
	if name == "" {
		return messaging.Content{}, messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v", name)))
	}
	ct, err := a.cache.get(name)
	if err == nil {
		return ct, messaging.StatusOK()
	}
	ct2, status := a.intf.Representation(http.MethodGet, name, "", "", nil)
	if !status.OK() {
		return messaging.Content{}, status
	}
	a.cache.put(name, ct2)
	return messaging.Content{}, messaging.StatusNotFound()
}

func (a *agentT) putRepresentation(name, author, contentType string, value any) *messaging.Status {
	if name == "" || author == "" || contentType == "" || value == nil {
		return messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v", name)))
	}
	ct := messaging.Content{Type: contentType, Value: value}
	buf, status := messaging.Marshal[[]byte](&ct)
	if !status.OK() {
		return status.WithLocation(name)
	}
	_, status2 := a.intf.Representation(http.MethodPut, name, author, contentType, buf)
	if !status2.OK() {
		return status.WithLocation(name)
	}
	// TODO: remove after initial testing
	a.cache.put(name, ct)
	return status
}

/*
	if nsName == "" {
		err = errors.New(fmt.Sprintf("nsName is empty on call to PutValue()"))
		return messaging.NewStatusError(http.StatusBadRequest, err, r.agent.Uri())
	}
	if resource == nil {
		err = errors.New(fmt.Sprintf("resource is nil on call to PutValue() for nsName : %v", nsName))
		return messaging.NewStatusError(http.StatusNoContent, err, r.agent.Uri())
	}

*/
/*
	var buf []byte
	var err error
	var ct string

	switch ptr := value.(type) {
	case string:
		ct = httpx.ContentTypeText
		buf = []byte(ptr)
		//v := text{ptr}
		//buf, err = json.Marshal(v)
		//if err != nil {
		//	return messaging.NewStatus(messaging.StatusJsonEncodeError, err)
		//}
	case []byte:
		ct = httpx.ContentTypeBinary
		buf = ptr
	case *url.URL:
		buf, err = iox.ReadFile(ptr)
		if err != nil {
			return messaging.NewStatus(messaging.StatusIOError, err)
		}
	default:
		buf, err = json.Marshal(ptr)
		if err != nil {
			return messaging.NewStatus(messaging.StatusJsonEncodeError, err)
		}
		ct = httpx.ContentTypeJson
	}
	if len(buf) == 0 {
		err = errors.New(fmt.Sprintf("resource is empty on call to PutValue() for nsName : %v", name))
		return messaging.NewStatus(http.StatusNoContent, err)
	}

*/
//_, status := httpPutContent(name, fragment, author, ct, buf)
//if !status.OK() {
//	return status.WithMessage(fmt.Sprintf("name %v", name))
//}
