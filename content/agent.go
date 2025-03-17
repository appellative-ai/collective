package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/io"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"net/url"
	"time"
)

const (
	AgentNamespaceName = "resiliency:agent/behavioral-ai/collective/content"
	agentUri           = AgentNamespaceName
	defaultDuration    = time.Second * 10
)

type agentT struct {
	running  bool
	agentId  string
	hostName string
	uri      []string
	duration time.Duration
	cache    *contentT
	mapCache *mapT

	ticker     *messaging.Ticker
	emissary   *messaging.Channel
	master     *messaging.Channel
	notifier   messaging.NotifyFunc
	dispatcher messaging.Dispatcher
	activity   messaging.ActivityFunc
}

func newAgent(dispatcher messaging.Dispatcher) *agentT {
	a := new(agentT)
	a.agentId = agentUri
	a.duration = defaultDuration
	a.cache = newContentCache()
	a.mapCache = newMapCache()
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
		if m.ContentType() == messaging.ContentTypeNotify || m.ContentType() == messaging.ContentTypeActivity {
			return
		}
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

func (a *agentT) getValue(name string, version int) (buf []byte, status *messaging.Status) {
	if name == "" || version <= 0 {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), a.Uri())
	}
	var err error
	buf, err = a.cache.get(name, version)
	if err == nil {
		return buf, messaging.StatusOK()
	}
	// Cache miss
	buf, status = httpGetContent(name, version)
	if !status.OK() {
		status.SetAgent(a.Uri())
		status.SetMessage(fmt.Sprintf("name %v and version %v", name, version))
		return nil, status
	}
	a.cache.put(name, buf, version)
	return buf, messaging.StatusOK()
}

func (a *agentT) addValue(name, author string, content any, version int) *messaging.Status {
	if name == "" || author == "" || content == nil || version <= 0 {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument name %v version %v", name, version)), a.Uri())
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

	switch ptr := content.(type) {
	case string:
		v := text{ptr}
		buf, err = json.Marshal(v)
		if err != nil {
			return messaging.NewStatusError(messaging.StatusJsonEncodeError, err, a.Uri())
		}
	case []byte:
		buf = ptr
	case *url.URL:
		buf, err = io.ReadFile(ptr)
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
	_, status := httpPutContent(name, author, buf, version)
	if !status.OK() {
		status.SetAgent(a.Uri())
		status.SetMessage(fmt.Sprintf("name %v and version %v", name, version))
		return status
	}
	a.cache.put(name, buf, version)
	return status
}

func (a *agentT) getAttributes(name string) (map[string]string, *messaging.Status) {
	if name == "" {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("map name [%v] is empty", name)), a.Uri())
	}
	m, err := a.mapCache.get(name)
	if err == nil {
		return m, messaging.StatusOK()
	}
	// Cache miss
	buf, status := httpGetContent(name, 1)
	if !status.OK() {
		status.SetAgent(a.Uri())
		status.SetMessage(fmt.Sprintf("map name [%v] not found", name))
		return nil, status
	}
	// TODO : parse buf into map
	if len(buf) > 0 {
	}
	return nil, messaging.StatusNotFound().SetAgent(a.Uri())
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
	return messaging.StatusOK() //BadRequest().SetAgent(a.Uri())
}
