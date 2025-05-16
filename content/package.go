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

// Representation -
type Representation struct {
	Version string // returned on a Get
	Type    string // Content-Type
	Content any
}

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
	Get func(collective, name, resource string) (Accessor, *messaging.Status)
	Add func(name, resource, author, authority string, access Accessor) *messaging.Status
	//List func(name string) ([]string, *messaging.Status)

	Representation    func(collective, name, fragment string) (Representation, *messaging.Status)
	AddRepresentation func(name, fragment, author string, rep Representation) *messaging.Status

	Context    func(collective, name, fragment string) (Representation, *messaging.Status)
	AddContext func(name, fragment, author string, rep Representation) *messaging.Status
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Get: func(collective, name, resource string) (Accessor, *messaging.Status) {
			return agent.getContent(name, resource)
		},
		Add: func(name, resource, author, authority string, access Accessor) *messaging.Status {
			// TODO: add collective name
			return agent.addContent(name, resource, author, authority, access)
		},

		Representation: func(collective, name, resource string) (Representation, *messaging.Status) {
			return Representation{}, messaging.StatusOK()
		},
		AddRepresentation: func(name, fragment, author string, rep Representation) *messaging.Status {
			// TODO: add collective name
			return messaging.StatusOK()
		},

		//List: func(name string) ([]string, *messaging.Status) {
		//	return nil, nil
		//},
	}
}()

// Resolve - generic typed resolution
// TODO: support map[string]string??
func Resolve[T any](collective, name, resource string, resolver *Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", name)))
	}
	access, status := resolver.Get(collective, name, resource)
	if !status.OK() {
		return t, status
	}
	if access.Content == nil {
		return t, messaging.NewStatus(http.StatusNoContent, fmt.Sprintf("content not found for name: %v", name))
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
				return t, messaging.NewStatus(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, name)))
			}
		}
	}
	return t, messaging.StatusOK()
}
