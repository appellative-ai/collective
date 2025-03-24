package eventing

import "fmt"

func ExampleActivity() {
	a := ActivityItem{
		Agent:   nil,
		Event:   "eventing",
		Source:  "source",
		Content: nil,
	}
	Activity(a)

	//fmt.Printf("test: Activity() -> [%v]\n",Activity(a))

	//Output:
	//active-> 2025-03-17T19:29:39.162Z [<nil>] [eventing] [source] [<nil>]

}

func ExampleActivityMessage() {
	//status := NewStatusMessage(http.StatusTeapot, "test message", "agent/test")
	m := NewActivityMessage(ActivityItem{
		Agent:   nil,
		Event:   "eventing",
		Source:  "source",
		Content: nil,
	})
	e := ActivityContent(m)
	fmt.Printf("test: ActivityContent() -> [%v]\n", e)

	//Output:
	//test: ActivityContent() -> [{<nil> eventing source <nil>}]

}
