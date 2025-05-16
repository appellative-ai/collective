package test

import (
	"fmt"
)

const (
	ResiliencyThreshold = "behavioral-ai:resiliency:type/operative/threshold"
	ResiliencyInterpret = "behavioral-ai:resiliency:type/operative/interpret"
)

// Testing
//r.activity = func(hostName string, agent messaging.Agent, eventing, source string, resource any) {
//	fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", messaging.FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), eventing, source, resource)
//}

func Startup() {
	status := loadResolver(resource.Resolver)
	if !status.OK() {
		fmt.Printf("error on loading Resolver: %v\n", status)
	}
	status = LoadProfile(resource.Resolver)
	if !status.OK() {
		fmt.Printf("error on loading Resolver: %v\n", status)
	}
}
