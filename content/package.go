package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

// HttpExchange - exchange type
type HttpExchange func(r *http.Request) (*http.Response, error)

// Resolution - in the real world
type Resolution interface {
	GetValue(name string, version int) ([]byte, *messaging.Status)
	PutValue(name, author string, content any, version int) *messaging.Status
	GetAttributes(name string) (map[string]string, *messaging.Status)
	PutAttributes(name, author string, m map[string]string) *messaging.Status
	AddActivity(agent messaging.Agent, event, source string, content any)
	Notify(e messaging.Event)
}

// Resolver - content resolution in the real world
var (
	Resolver = newHttpResolver()
)

func init() {
	if r, ok := any(Resolver).(*resolution); ok {
		// Testing
		r.notifier = messaging.Notify
		r.agent.notifier = r.notifier
		r.activity = func(hostName string, agent messaging.Agent, event, source string, content any) {
			fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
		}
	}
}

// Startup - run the agents
func Startup(uri []string, do HttpExchange, hostName string) {
	if r, ok := any(Resolver).(*resolution); ok {
		if do != nil {
			r.do = do
		}
		r.agent.uri = uri
		r.agent.hostName = hostName
		r.agent.Run()
	}
}

func Shutdown() {
	if r, ok := any(Resolver).(*resolution); ok {
		r.agent.Shutdown()
	}
}

// Resolve - generic typed resolution
func Resolve[T any](name string, version int, resolver Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", name)), Name)
	}
	body, status := resolver.GetValue(name, version)
	if !status.OK() {
		return t, status
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, status1 := Resolve[text](name, version, resolver)
		if !status1.OK() {
			return t, status1
		}
		*ptr = t1.Value
	case *[]byte:
		*ptr = body
	default:
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, name)), Name)
		}
	}
	return t, messaging.StatusOK()
}

// NewEphemeralResolver - in memory resolver
func NewEphemeralResolver() Resolution {
	return initializedEphemeralResolver(true, true)
}

// NewConfigEphemeralResolver - in memory resolver
func NewConfigEphemeralResolver(activity, notify bool) Resolution {
	return initializedEphemeralResolver(activity, notify)
}
