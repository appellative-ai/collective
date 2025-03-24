package operations

import (
	"errors"
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleNewAgent() {
	agent := newAgent()
	status := messaging.NewStatusError(http.StatusTeapot, errors.New("error"), agent.Uri())
	status.WithMessage("notify message")
	status.WithRequestId("123-request-id")
	agent.Message(eventing.NewNotifyMessage(status))

	//Output:
	//fail

}
