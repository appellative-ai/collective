package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// Common resources

const (
	ListResource     = "ls"   // list for a name, allow arguments for filtering
	SrcResource      = "src"  // location of source code, optional versioning
	InstanceResource = "inst" // the state/instance of a type, optional versioning
	InfoResource     = "info" // information, help
	InfoResourceAlt  = "?"
)

// Accessor -
type Accessor struct {
	Version string // returned on a Get
	Type    string // Content-Type
	Content any
}

func (a Accessor) String() string {
	return fmt.Sprintf("vers: %v type: %v content: %v", a.Version, a.Type, a.Content != nil)
}

// Resolution - in the real world
// Can only add in current collective. An empty collective is assuming the local vs distributed
// How to handle local vs distributed
type Resolution struct {
	Get func(collective, name, resource string) (Accessor, *Status)
	Add func(name, resource, author, authority string, access Accessor) *Status
	//List func(name string) ([]string, *messaging.Status)
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Get: func(collective, name, resource string) (Accessor, *Status) {
			return agent.getContent(name, resource)
		},
		Add: func(name, resource, author, authority string, access Accessor) *Status {
			// TODO: add collective name
			return agent.addContent(name, resource, author, authority, access)
		},
		//List: func(name string) ([]string, *messaging.Status) {
		//	return nil, nil
		//},
	}
}()

// Resolve - generic typed resolution
// TODO: support map[string]string??
func Resolve[T any](collective, name, resource string, resolver *Resolution) (T, *Status) {
	var t T

	if resolver == nil {
		return t, NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", name)))
	}
	access, status := resolver.Get(collective, name, resource)
	if !status.OK() {
		return t, status
	}
	if access.Content == nil {
		return t, NewStatus(http.StatusNoContent, fmt.Sprintf("content not found for name: %v", name))
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, status1 := Resolve[text](collective, name, resource, resolver)
		if !status1.OK() {
			return t, status1
		}
		*ptr = t1.Value
	case *[]byte:
		if body, ok := access.Content.([]byte); ok {
			*ptr = body
		}
	default:
		var (
			body []byte
			ok   bool
		)
		body, ok = access.Content.([]byte)
		if ok {
			err := json.Unmarshal(body, ptr)
			if err != nil {
				return t, NewStatus(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, name)))
			}
		}
	}
	return t, StatusOK()
}
