package namespace

import (
	"github.com/behavioral-ai/core/messaging"
)

// Appender - append
type Appender struct {
	Thing    func(nsName, author string) *messaging.Status
	Relation func(nsName1, nsName2, author string) *messaging.Status
}

// Append -
var Append = func() *Appender {
	return &Appender{
		Thing: func(nsName, author string) *messaging.Status {
			return messaging.StatusBadRequest()
		},
		Relation: func(nsName1, nsName2, author string) *messaging.Status {
			return messaging.StatusBadRequest()
		},
	}
}()
