package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"net/url"
	"time"
)

const (
	namespaceName     = "core:agent/collective/resource"
	defaultDuration   = time.Second * 10
	contentTypeBinary = "application/octet-stream"
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

	//handler  eventing.Agent
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

func NewAgent() messaging.Agent {
	agent = newAgent()
	return agent //host.Register(agent)
}

func newAgent() *agentT {
	a := new(agentT)
	a.duration = defaultDuration
	a.cache = newCache()

	//a.handler = handler
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return namespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if !a.running {
		if m.Name() == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Name() == messaging.StartupEvent {
			a.run()
			a.running = true
			return
		}
		return
	}
	if m.Name() == messaging.ShutdownEvent {
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
		messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Name())
	}
	// configure
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

func (a *agentT) getRepresentation(name, fragment string) (Content, *messaging.Status) {
	if name == "" {
		return Content{}, messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v", name)))
	}
	ct, err := a.cache.get(name, fragment)
	if err == nil {
		return ct, messaging.StatusOK()
	}
	// Cache miss
	buf, err1 := httpGetContent(name)
	if err1 != nil {
		return Content{}, err1.WithMessage(fmt.Sprintf("name %v", name))
	}
	ct = Content{Fragment: fragment, Type: http.DetectContentType(buf), Value: buf}
	a.cache.put(name, fragment, ct)
	return ct, messaging.StatusOK()
}

func (a *agentT) putRepresentation(name, fragment, author string, ct Content) *messaging.Status {
	if name == "" || author == "" || ct.Value == nil {
		return messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v", name)))
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
	var buf []byte
	var err error

	switch ptr := ct.Value.(type) {
	case string:
		buf = []byte(ptr)
		//v := text{ptr}
		//buf, err = json.Marshal(v)
		//if err != nil {
		//	return messaging.NewStatus(messaging.StatusJsonEncodeError, err)
		//}
		//ct.Type = httpx.ContentTypeText
		//ct.Value = buf
	case []byte:
		buf = ptr
		//ct.Type = contentTypeBinary
		//ct.Value = buf
	case map[string]string:
		s := iox.WriteMap(ptr)
		buf = []byte(s)
		ct.Type = httpx.ContentTypeText
		ct.Value = s
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
		ct.Type = httpx.ContentTypeJson
		ct.Value = string(buf)
	}
	if len(buf) == 0 {
		err = errors.New(fmt.Sprintf("resource is empty on call to PutValue() for nsName : %v", name))
		return messaging.NewStatus(http.StatusNoContent, err)
	}
	_, status := httpPutContent(name, fragment, author, ct.Type, buf)
	if !status.OK() {
		return status.WithMessage(fmt.Sprintf("name %v", name))
	}
	a.cache.put(name, fragment, ct) //Accessor{Version: uri.UnnVersion(name), Type: http.DetectContentType(buf), Content: buf})
	return status
}
