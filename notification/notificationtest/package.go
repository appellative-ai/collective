package notificationtest

import (
	"fmt"
	"github.com/appellative-ai/collective/notification"
	"github.com/appellative-ai/core/fmtx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
	"time"
)

// NewNotifier -
func NewNotifier() *notification.Interface {
	return &notification.Interface{
		Message: func(msg *messaging.Message) *std.Status {
			fmt.Printf("%v  -> %v\n", "message", msg)
			return std.StatusOK
		},
		Trace: func(name, task, observation, action string) {
			fmt.Printf("%v [%v] [%v] [%v] [%v]", fmtx.FmtRFC3339Millis(time.Now().UTC()), name, task, observation, action)
		},
	}
}

/*
// NewReceiver -
func NewReceiver() *notification.Receiver {
	return &notification.Receiver{
		Message: func(name string) *messaging.Message {
			//fmt.Printf("%v  -> %v\n", "message", msg)
			return nil
		},
		Advice: func(name string) *messaging.Message {
			return nil
		},
	}
}


*/
