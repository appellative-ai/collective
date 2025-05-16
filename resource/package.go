package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
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
	Representation    func(collective, name, fragment string) (Content, *messaging.Status)
	AddRepresentation func(name, fragment, author string, ct Content) *messaging.Status

	Context    func(collective, name string) (Content, *messaging.Status)
	AddContext func(name, author string, ct Content) *messaging.Status
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Representation: func(collective, name, fragment string) (Content, *messaging.Status) {
			return Content{}, messaging.StatusOK()
		},
		AddRepresentation: func(name, fragment, author string, ct Content) *messaging.Status {
			// TODO: add collective name
			return messaging.StatusOK()
		},

		Context: func(collective, name string) (Content, *messaging.Status) {
			return Content{}, messaging.StatusOK()
		},
		AddContext: func(name, author string, ct Content) *messaging.Status {
			// TODO: add collective name
			return messaging.StatusOK()
		},
	}
}()

// Resolve - generic typed resolution
// TODO: support map[string]string??
func Resolve[T any](collective, name, fragment string, resolver *Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", name)))
	}
	ct, status := resolver.Representation(collective, name, fragment)
	if !status.OK() {
		return t, status
	}
	if ct.Value == nil {
		return t, messaging.NewStatus(http.StatusNoContent, fmt.Sprintf("resource not found for name: %v", name))
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, status1 := Resolve[text](collective, name, fragment, resolver)
		if !status1.OK() {
			return t, status1
		}
		*ptr = t1.Value
	case *[]byte:
		if body, ok := ct.Value.([]byte); ok {
			*ptr = body
		}
	default:
		var (
			body []byte
			ok   bool
		)
		body, ok = ct.Value.([]byte)
		if ok {
			err := json.Unmarshal(body, ptr)
			if err != nil {
				return t, messaging.NewStatus(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, name)))
			}
		}
	}
	return t, messaging.StatusOK()
}
