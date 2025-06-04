package resource

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// Resolution - in the real world
// Can only add in current collective. An empty collective is assuming the local vs distributed
// How to handle local vs distributed
type Resolution struct {
	Representation    func(name, fragment string) (messaging.Content, *messaging.Status)
	AddRepresentation func(name, author string, ct messaging.Content) *messaging.Status

	Context    func(name string) (messaging.Content, *messaging.Status)
	AddContext func(name, author string, ct messaging.Content) *messaging.Status
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Representation: func(name, fragment string) (messaging.Content, *messaging.Status) {
			return agent.getRepresentation(name, fragment)
		},
		AddRepresentation: func(name, author string, ct messaging.Content) *messaging.Status {
			return agent.putRepresentation(name, author, ct)
		},
		Context: func(name string) (messaging.Content, *messaging.Status) {
			return messaging.Content{}, messaging.StatusOK()
		},
		AddContext: func(name, author string, ct messaging.Content) *messaging.Status {
			return messaging.StatusOK()
		},
	}
}()

// Resolve - generic typed resolution
// TODO: test augment result from unmarshall with name
func Resolve[T any](name, fragment string, resolver *Resolution) (t T, status *messaging.Status) {
	if resolver == nil {
		return t, messaging.NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: resolver is nil for : %v", name)))
	}
	ct, status1 := resolver.Representation(name, fragment)
	if !status1.OK() {
		return t, status1
	}
	t, status = messaging.Unmarshal[T](&ct)
	if !status.OK() {
		status.WithLocation(name)
	}
	return
}
