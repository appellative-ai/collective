package eventtest

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleNewAgent() {
	a := newAgent(nil)

	status := messaging.NewStatusError(http.StatusTeapot, errors.New("error message"), a.Uri())
	a.Message(eventing.NewNotifyMessage(status))
	a.Message(eventing.NewDispatchMessage(a, messaging.NewChannel("test"), "event:teapot"))
	a.Message(eventing.NewActivityMessage(eventing.ActivityItem{
		Agent:   a,
		Event:   "activity-event",
		Source:  "source",
		Content: nil,
	}))

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [resiliency:agent/behavioral-ai/collective/eventing]

}
