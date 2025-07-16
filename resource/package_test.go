package resource

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
)

type Address struct {
	Line1 string
	Line2 string
	City  string
	State string
	Zip   string
}

func ExampleResolveBinary() {
	NewAgent()
	name := "core:type/binary"
	bytes := []byte("this is a test buffer")

	status := agent.putRepresentation(name, "author", messaging.ContentTypeBinary, bytes)
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", string(buf), status2)
	}

	s3, status3 := messaging.New[[]byte](&ct)
	fmt.Printf("test: New[[]byte]() -> [value:%v] [status:%v]\n", string(s3), status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment:  type: application/octet-stream value: true] [status:OK]
	//test: Representation() -> [value:this is a test buffer] [status:OK]
	//test: New[[]byte]() -> [value:this is a test buffer] [status:OK]

}

func ExampleResolveString() {
	NewAgent()
	name := "core:type/binary"
	s := "this is a test string"

	status := agent.putRepresentation(name, "author", messaging.ContentTypeText, s)
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", string(buf), status2)
	}

	s3, status3 := messaging.New[string](&ct)
	fmt.Printf("test: New[string]() -> [value:%v] [status:%v]\n", string(s3), status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment:  type: text/plain charset=utf-8 value: true] [status:OK]
	//test: New[string]() -> [value:this is a test string] [status:OK]

}

func ExampleResolveType() {
	NewAgent()
	addr := Address{
		Line1: "123 Main",
		Line2: "",
		City:  "Anytown",
		State: "Ohio",
		Zip:   "54321",
	}
	name := "core:type/address"

	status := agent.putRepresentation(name, "author", messaging.ContentTypeJson, addr)
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
	}

	s3, status3 := messaging.New[Address](&ct)
	fmt.Printf("test: New[Address]() -> [value:%v] [status:%v]\n", s3, status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment:  type: application/json value: true] [status:OK]
	//test: New[Address]() -> [value:{123 Main  Anytown Ohio 54321}] [status:OK]

}

func ExampleResolveMap() {
	NewAgent()
	m := map[string]string{
		"Line1": "123 Main",
		"Line2": "",
		"City":  "Anytown",
		"State": "Ohio",
		"Zip":   "54321",
	}
	name := "core:type/map"
	//fragment := "v2"

	status := agent.putRepresentation(name, "author", messaging.ContentTypeJson, m)
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
	}

	s3, status3 := messaging.New[map[string]string](&ct)
	fmt.Printf("test: New[map[string]string]() -> [value:%v] [status:%v]\n", s3, status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment:  type: application/json value: true] [status:OK]
	//test: New[map[string]string]() -> [value:map[City:Anytown Line1:123 Main Line2: State:Ohio Zip:54321]] [status:OK]

}
