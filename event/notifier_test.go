package event

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func ExampleNotify() {
	s := messaging.NewStatusError(http.StatusGatewayTimeout, errors.New("rate limiting"), "test:agent")
	s.WithAgent("resiliency:agent/operative")
	s.WithRequestId("123-request-id")

	Notify(s)

	//Output:
	//notify-> 2025-02-26T15:34:45.784Z [resiliency:agent/operative] [core:messaging.status] [123-request-id] [Timeout] [rate limiting]

}

func ExampleNotifyMessage() {
	status := messaging.NewStatusWithMessage(http.StatusTeapot, "test message", "agent/test")
	m := NewNotifyMessage(status)
	e := NotifyContent(m)
	fmt.Printf("test: NotifyContent() -> [%v]\n", e)

	//Output:
	//test: NotifyContent() -> [I'm A Teapot [msg:test message] [agent:agent/test]]

}

func ExampleNewStatusError() {
	s := messaging.NewStatusError(http.StatusGatewayTimeout, errors.New("rate limited"), "test:agent") //"resiliency:agent/operative/agent1#us-west")
	fmt.Printf("test: NewStatusError() -> [%v]\n", s)

	if _, ok := any(s).(NotifyItem); ok {
		fmt.Printf("test: Event() -> [%v]\n", ok)

	}

	//Output:
	//test: NewStatusError() -> [Timeout [err:rate limited] [agent:test:agent]]
	//test: Event() -> [true]

}

func ExampleNewStatusMessage() {
	s := messaging.NewStatusWithMessage(http.StatusOK, "successfully change ticker duration", "test:agent")
	fmt.Printf("test: NewStatusMessage() -> [%v]\n", s)

	if _, ok := any(s).(NotifyItem); ok {
		fmt.Printf("test: Event() -> [%v]\n", ok)

	}

	//Output:
	//test: NewStatusMessage() -> [OK [msg:successfully change ticker duration] [agent:test:agent]]
	//test: Event() -> [true]

}
