package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/collective/operations"
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/uri"
	"net/http"
	"net/url"
	"time"
)

const (
	NamespaceName   = "unn:behavioral-ai.github.com:resiliency:agent/collective/content"
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
	cache    *contentT

	handler  eventing.Agent
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

func init() {
	agent = newAgent(eventing.Handler)
	operations.Register(agent)
}

func newAgent(handler eventing.Agent) *agentT {
	a := new(agentT)
	a.duration = defaultDuration
	a.cache = newContentCache()

	a.handler = handler
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if !a.running {
		if m.Event() == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Event() == messaging.StartupEvent {
			a.run()
			a.running = true
			return
		}
		return
	}
	if m.Event() == messaging.ShutdownEvent {
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
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
	}
	// configure
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
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

func (a *agentT) getValue(name, resource string) (access Accessor, status *messaging.Status) {
	if name == "" {
		return Accessor{}, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v", name)), a.Uri())
	}
	var err error
	access, err = a.cache.get(name, resource)
	if err == nil {
		return access, messaging.StatusOK()
	}
	// Cache miss
	var buf []byte
	buf, status = httpGetContent(name)
	if !status.OK() {
		status.WithAgent(a.Uri())
		status.WithMessage(fmt.Sprintf("name %v", name))
		return access, status
	}
	a.cache.put(name, resource, Accessor{Version: uri.UnnVersion(name), Type: http.DetectContentType(buf), Content: buf})
	return access, messaging.StatusOK()
}

func (a *agentT) addValue(name, resource, author, authority string, access Accessor) *messaging.Status {
	if name == "" || author == "" || authority == "" || access.Content == nil {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v", name)), a.Uri())
	}
	/*
		if nsName == "" {
			err = errors.New(fmt.Sprintf("nsName is empty on call to PutValue()"))
			return messaging.NewStatusError(http.StatusBadRequest, err, r.agent.Uri())
		}
		if content == nil {
			err = errors.New(fmt.Sprintf("content is nil on call to PutValue() for nsName : %v", nsName))
			return messaging.NewStatusError(http.StatusNoContent, err, r.agent.Uri())
		}

	*/
	var buf []byte
	var err error

	switch ptr := access.Content.(type) {
	case string:
		v := text{ptr}
		buf, err = json.Marshal(v)
		if err != nil {
			return messaging.NewStatusError(messaging.StatusJsonEncodeError, err, a.Uri())
		}
	case []byte:
		buf = ptr
	case *url.URL:
		buf, err = iox.ReadFile(ptr)
		if err != nil {
			return messaging.NewStatusError(messaging.StatusIOError, err, a.Uri())
		}
	default:
		buf, err = json.Marshal(ptr)
		if err != nil {
			return messaging.NewStatusError(messaging.StatusJsonEncodeError, err, a.Uri())
		}
	}
	if len(buf) == 0 {
		err = errors.New(fmt.Sprintf("content is empty on call to PutValue() for nsName : %v", name))
		return messaging.NewStatusError(http.StatusNoContent, err, a.Uri())
	}
	_, status := httpPutContent(name, authority, author, buf)
	if !status.OK() {
		status.WithAgent(a.Uri())
		status.WithMessage(fmt.Sprintf("name %v", name))
		return status
	}
	a.cache.put(name, resource, access) //Accessor{Version: uri.UnnVersion(name), Type: http.DetectContentType(buf), Content: buf})
	return status
}

/*
func (a *agentT) getAttributes(name string) (map[string]string, *messaging.Status) {
	if name == "" {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("map name [%v] is empty", name)), a.Uri())
	}
	m, err := a.mapCache.get(name)
	if err == nil {
		return m, messaging.StatusOK()
	}
	// Cache miss
	buf, status := httpGetContent(name, resource,version)
	if !status.OK() {
		status.WithAgent(a.Uri())
		status.WithMessage(fmt.Sprintf("map name [%v] not found", name))
		return nil, status
	}
	// TODO : parse buf into map
	if len(buf) > 0 {
	}
	return nil, messaging.StatusNotFound().WithAgent(a.Uri())
}

func (a *agentT) addAttributes(name, author string, m map[string]string) *messaging.Status {
	if name == "" || author == "" || m == nil {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid argument name [%v],author [%v] or map", name, author)), a.Uri())
	}
	err := a.mapCache.put(name, m)
	if err == nil {
		return messaging.StatusOK()
	}
	//buf,status := httpPutContent(name,author,)
	return messaging.StatusOK() //BadRequest().WithAgent(a.Uri())
}


*/
