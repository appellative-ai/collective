package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"reflect"
)

// Content -
type Content struct {
	Fragment string // returned on a Get
	Type     string // Content-Type
	Value    any
}

func (c Content) String() string {
	return fmt.Sprintf("fragment: %v type: %v value: %v", c.Fragment, c.Type, c.Value != nil)
}

// Resolution - in the real world
// Can only add in current collective. An empty collective is assuming the local vs distributed
// How to handle local vs distributed
type Resolution struct {
	Representation    func(name, fragment string) (Content, *messaging.Status)
	AddRepresentation func(name, fragment, author string, value any) *messaging.Status

	Context    func(name string) (Content, *messaging.Status)
	AddContext func(name, author string, ct Content) *messaging.Status

	// TODO: Need some sort of context and then the result, what was the member working on, and what was the
	// result of that work, with pertinent details. Do not need location
	// Should be able to monitor cause and effect
	AddTrace func(name, task, observation, action string) *messaging.Status
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Representation: func(name, fragment string) (Content, *messaging.Status) {
			return agent.getRepresentation(name, fragment)
		},
		AddRepresentation: func(name, fragment, author string, value any) *messaging.Status {
			return agent.putRepresentation(name, fragment, author, value)
		},
		Context: func(name string) (Content, *messaging.Status) {
			return Content{}, messaging.StatusOK()
		},
		AddContext: func(name, author string, ct Content) *messaging.Status {
			return messaging.StatusOK()
		},
		AddTrace: func(name, task, observation, action string) *messaging.Status {
			return messaging.StatusOK()
		},
	}
}()

// Resolve - generic typed resolution
// TODO: support map[string]string??
func Resolve[T any](name, fragment string, resolver *Resolution) (T, *messaging.Status) {
	var t T
	var body []byte
	var ok bool

	if resolver == nil {
		return t, messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", name)))
	}
	ct, status := resolver.Representation(name, fragment)
	if !status.OK() {
		return t, status
	}
	if ct.Value == nil {
		return t, messaging.NewStatus(http.StatusNoContent, fmt.Sprintf("representation not found for name: %v", name))
	}
	if body, ok = ct.Value.([]byte); !ok {
		return t, messaging.NewStatus(messaging.StatusInvalidContent, fmt.Sprintf("representation content type is not []byte for name: %v", name))
	}
	switch ptr := any(&t).(type) {
	case *string:
		if ct.Type != httpx.ContentTypeText {
			return t, messaging.NewStatus(messaging.StatusInvalidContent, fmt.Sprintf("representation content type %v invalid for string: %v", ct.Type, name))
		}
		*ptr = string(body)
	case *[]byte:
		if ct.Type != httpx.ContentTypeBinary {
			return t, messaging.NewStatus(messaging.StatusInvalidContent, fmt.Sprintf("representation content type %v invalid for []byte: %v", ct.Type, name))
		}
		*ptr = body
	default:
		if ct.Type != httpx.ContentTypeJson {
			return t, messaging.NewStatus(messaging.StatusInvalidContent, fmt.Sprintf("representation content type %v invalid for %v: %v", ct.Type, reflect.TypeOf(t), name))
		}
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, messaging.NewStatus(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, name)))
		}
	}
	return t, messaging.StatusOK()
}
