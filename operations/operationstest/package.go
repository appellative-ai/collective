package operationstest

import (
	"fmt"
	"github.com/behavioral-ai/collective/operations"
	"github.com/behavioral-ai/core/fmtx"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

// NewService -
func NewService() *operations.Service {
	return &operations.Service{
		Message: func(msg *messaging.Message) {
			fmt.Printf("%v  -> %v\n", "message", msg)
		},
		Advise: func(msg *messaging.Message) {
			fmt.Printf("%v   -> %v\n", "advise", msg)
		},
		Subscribe: func(msg *messaging.Message) {
			fmt.Printf("%v-> %v\n", "subscribe", msg)
		},
		CancelSubscription: func(msg *messaging.Message) {
			fmt.Printf("%v   -> %v\n", "cancel", msg)
		},
		Trace: func(name, task, observation, action string) {
			fmt.Printf("%v [%v] [%v] [%v] [%v]", fmtx.FmtRFC3339Millis(time.Now().UTC()), name, task, observation, action)
		},
	}
}
