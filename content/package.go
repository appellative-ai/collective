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
type Resolution struct {
	Get func(name string) (Accessor, *messaging.Status)
	Add func(name, author string, content any) *messaging.Status
	//List func(name string) ([]string, *messaging.Status)
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Get: func(name string) (Accessor, *messaging.Status) {
			return agent.getValue(name)
		},
		Add: func(name, author string, content any) *messaging.Status {
			return agent.addValue(name, author, content)
		},
		//List: func(name string) ([]string, *messaging.Status) {
		//	return nil, nil
		//},
	}
}()

// Resolve - generic typed resolution
// TODO: support map[string]string??
func Resolve[T any](name string, resolver *Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", name)), NamespaceName)
	}
	access, status := resolver.Get(name)
	if !status.OK() {
		return t, status
	}
	if access.Content == nil {
		return t, messaging.NewStatusWithMessage(http.StatusNoContent, fmt.Sprintf("content not found for name: %v", name), NamespaceName)
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, status1 := Resolve[text](name, resolver)
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
				return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, name)), NamespaceName)
			}
		}
	}
	return t, messaging.StatusOK()
}
