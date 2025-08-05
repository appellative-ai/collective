package resolution

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
	"net/http"
	"time"
)

const (
	AgentName = "common:core:agent/resolution/collective"
	duration  = time.Second * 10
	timeout   = time.Second * 3
)

var (
	agent *agentT
)

type text struct {
	Value string
}

type agentT struct {
	running bool
	timeout time.Duration
	cache   *cacheT

	ex       rest.Exchange
	logFunc  func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)
	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
}

func NewAgent() messaging.Agent {
	return newAgent()
}

func newAgent() *agentT {
	a := new(agentT)
	agent = a
	a.timeout = timeout
	a.cache = newCache()

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return AgentName }

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
		messaging.UpdateContent[time.Duration](m, &a.timeout)
		messaging.UpdateContent[func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)](m, &a.logFunc)
		return
	case messaging.StartupEvent:
		if a.running {
			return
		}
		a.running = true
		a.run()
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

// Run - run the agent
func (a *agentT) run() {
	go emissaryAttend(a)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) getRepresentation(name string) (std.Content, *std.Status) {
	if name == "" {
		return std.Content{}, std.NewStatus(http.StatusBadRequest, "", errors.New(fmt.Sprintf("error: invalid argument name %v", name)))
	}
	ct, err := a.cache.get(name)
	if err == nil {
		return ct, std.StatusOK
	}
	//ct2, status := a.intf.Representation(http.MethodGet, name, "", "", nil)
	//if !status.OK() {
	//	return std.Content{}, status
	//}
	//a.cache.put(name, ct2)
	return std.Content{}, std.StatusNotFound
}

func (a *agentT) putRepresentation(name, author, contentType string, value any) *std.Status {
	if name == "" || author == "" || contentType == "" || value == nil {
		return std.NewStatus(http.StatusBadRequest, "", errors.New(fmt.Sprintf("error: invalid argument name %v", name)))
	}
	ct := std.Content{Type: contentType, Value: value}
	_, status := std.Marshal[[]byte](&ct)
	if !status.OK() {
		return status //.WithLocation(name)
	}
	//_, status2 := a.intf.Representation(http.MethodPut, name, author, contentType, buf)
	//if !status2.OK() {
	//	return status //.WithLocation(name)
	//}
	// TODO: remove after initial testing
	//a.cache.put(name, ct)
	return status
}

/*
	if nsName == "" {
		err = errors.New(fmt.Sprintf("nsName is empty on call to PutValue()"))
		return messaging.NewStatusError(http.StatusBadRequest, err, r.agent.Uri())
	}
	if resolution == nil {
		err = errors.New(fmt.Sprintf("resolution is nil on call to PutValue() for nsName : %v", nsName))
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
		err = errors.New(fmt.Sprintf("resolution is empty on call to PutValue() for nsName : %v", name))
		return messaging.NewStatus(http.StatusNoContent, err)
	}

*/
//_, status := httpPutContent(name, fragment, author, ct, buf)
//if !status.OK() {
//	return status.WithMessage(fmt.Sprintf("name %v", name))
//}
