package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// Resolution1 - in the real world
type Resolution1 interface {
	GetValue(nsName string, version int) ([]byte, *messaging.Status)
	AddValue(nsName, author string, content any, version int) *messaging.Status
	GetAttributes(nsName string) (map[string]string, *messaging.Status)
	AddAttributes(nsName, author string, m map[string]string) *messaging.Status
}

// Resolution - in the real world
type Resolution struct {
	Get func(nsName string, version string) ([]byte, string, *messaging.Status)
	Add func(nsName, author, contentType string, content any, version string) *messaging.Status
	//GetAttributes func(nsName string) (map[string]string, *messaging.Status)
	//AddAttributes func(nsName, author string, m map[string]string) *messaging.Status
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Get: func(nsName, version string) ([]byte, string, *messaging.Status) {
			return agent.getValue(nsName, version)
		},
		Add: func(nsName, author, contentType string, content any, version string) *messaging.Status {
			return agent.addValue(nsName, author, contentType, content, version)
		},
		/*
			GetAttributes: func(nsName string) (map[string]string, *messaging.Status) {
				return agent.getAttributes(nsName)
			},
			AddAttributes: func(nsName, author string, m map[string]string) *messaging.Status {
				return agent.addAttributes(nsName, author, m)
			},

		*/
	}
}()

// Resolve - generic typed resolution
// TODO: support map[string]string??
func Resolve[T any](nsName, version string, resolver *Resolution) (T, string, *messaging.Status) {
	var t T

	if resolver == nil {
		return t, "", messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: BadRequest - resolver is nil for : %v", nsName)), NamespaceName)
	}
	body, ct, status := resolver.Get(nsName, version)
	if !status.OK() {
		return t, "", status
	}
	if len(body) == 0 {
		return t, ct, messaging.NewStatusWithMessage(http.StatusNoContent, fmt.Sprintf("content not found for name: %v", nsName), NamespaceName)
	}
	switch ptr := any(&t).(type) {
	case *string:
		t1, ct1, status1 := Resolve[text](nsName, version, resolver)
		if !status1.OK() {
			return t, ct1, status1
		}
		*ptr = t1.Value
	case *[]byte:
		*ptr = body
	default:
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, ct, messaging.NewStatusError(messaging.StatusJsonDecodeError, errors.New(fmt.Sprintf("JsonDecode - %v for : %v", err, nsName)), NamespaceName)
		}
	}
	return t, ct, messaging.StatusOK()
}
