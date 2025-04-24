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
	SelfResource     = "self" // maybe . information about the Unn, such as resource names
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

// Resolution - in the real world
type Resolution struct {
	Get       func(nsName, resource, version string) (Accessor, *messaging.Status)
	Add       func(nsName, resource, version, author string, content any) *messaging.Status
	Resources func(nsName string) ([]string, *messaging.Status)
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Get: func(nsName, resource, version string) (Accessor, *messaging.Status) {
			return agent.getValue(nsName, resource, version)
		},
		Add: func(nsName, resource, version, author string, content any) *messaging.Status {
			return agent.addValue(nsName, resource, version, author, content)
		},
		Resources: func(nsName string) ([]string, *messaging.Status) {
			return nil, nil
		},
	}
}()

// Resolve - generic typed resolution
// TODO: support map[string]string??
func Resolve[T any](nsName, resource, version string, resolver *Resolution) (T, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", nsName)), NamespaceName)
	}
	access, status := resolver.Get(nsName, resource, version)
	if !status.OK() {
		return t, status
	}
	if access.Content == nil {
		return t, messaging.NewStatusWithMessage(http.StatusNoContent, fmt.Sprintf("content not found for name: %v", nsName), NamespaceName)
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, status1 := Resolve[text](nsName, resource, version, resolver)
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
				return t, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, nsName)), NamespaceName)
			}
		}
	}
	return t, messaging.StatusOK()
}
