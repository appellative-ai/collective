package operationstest

import (
	"fmt"
	"github.com/appellative-ai/collective/operations"
	"github.com/appellative-ai/core/fmtx"
	"github.com/appellative-ai/core/messaging"
	"time"
)

// NewNotifier -
func NewNotifier() *operations.Notification {
	return &operations.Notification{
		Message: func(msg *messaging.Message) bool {
			fmt.Printf("%v  -> %v\n", "message", msg)
			return true
		},
		Advise: func(msg *messaging.Message) *messaging.Status {
			fmt.Printf("%v   -> %v\n", "advise", msg)
			return messaging.StatusOK()
		},
		Trace: func(name, task, observation, action string) {
			fmt.Printf("%v [%v] [%v] [%v] [%v]", fmtx.FmtRFC3339Millis(time.Now().UTC()), name, task, observation, action)
		},
	}
}
