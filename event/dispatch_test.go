package event

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func ExampleTraceDispatch_Channel() {
	d := NewFilteredTraceDispatcher(nil, messaging.Emissary)
	channel := ""
	event := ""

	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [channel:%v]\n", channel)

	channel = messaging.Emissary
	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [channel:%v]\n", channel)

	//Output:
	//test: Dispatch() -> [channel:]
	//trace -> 2024-11-24T18:40:08.606Z [<nil>] [emissary] []
	//test: Dispatch() -> [channel:emissary]

}

func ExampleTraceDispatch_Event() {
	d := NewFilteredTraceDispatcher([]string{messaging.ShutdownEvent, messaging.StartupEvent}, "")
	channel := ""
	event := ""

	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [%v]\n", event)

	event = messaging.ShutdownEvent
	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [%v]\n", event)

	event = messaging.ObservationEvent
	d.Dispatch(nil, channel, event)
	fmt.Printf("test: Dispatch() -> [channel:%v] [%v]\n", channel, event)

	//Output:
	//test: Dispatch() -> []
	//trace -> 2024-11-24T18:46:04.697Z [<nil>] [] [event:shutdown]
	//test: Dispatch() -> [event:shutdown]
	//test: Dispatch() -> [channel:] [event:observation]

}
