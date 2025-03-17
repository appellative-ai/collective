package namespace

import (
	http2 "github.com/behavioral-ai/core/http"
	"github.com/behavioral-ai/core/messaging"
)

var (
	Agent    messaging.Agent
	Exchange http2.Exchange
)

func init() {
	Exchange = http2.Do
}

// Adder - add
type Adder struct {
	Thing    func(nsName, author string) *messaging.Status
	Relation func(nsName1, nsName2, author string) *messaging.Status
}

// Add -
var Add = func() *Adder {
	return &Adder{
		Thing: func(nsName, author string) *messaging.Status {
			return messaging.StatusBadRequest()
		},
		Relation: func(nsName1, nsName2, author string) *messaging.Status {
			return messaging.StatusBadRequest()
		},
	}
}()
