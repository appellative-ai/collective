package resource

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
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
	fragment := "v1"

	status := Resolver.AddRepresentation(name, "author", messaging.Content{Fragment: fragment, Type: messaging.ContentTypeBinary, Value: bytes})
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name, fragment)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", string(buf), status2)
	}

	s3, status3 := Resolve[[]byte](name, fragment, Resolver)
	fmt.Printf("test: Resolve() -> [value:%v] [status:%v]\n", string(s3), status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment: v1 type: application/octet-stream value: true] [status:OK]
	//test: Representation() -> [value:this is a test buffer] [status:OK]
	//test: Resolve() -> [value:this is a test buffer] [status:OK]

}

func ExampleResolveString() {
	NewAgent()
	name := "core:type/binary"
	s := "this is a test string"
	fragment := "v1"

	status := Resolver.AddRepresentation(name, "author", messaging.Content{Fragment: fragment, Type: messaging.ContentTypeText, Value: s})
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name, fragment)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", string(buf), status2)
	}

	s3, status3 := Resolve[string](name, fragment, Resolver)
	fmt.Printf("test: Resolve() -> [value:%v] [status:%v]\n", string(s3), status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment: v1 type: text/plain charset=utf-8 value: true] [status:OK]
	//test: Representation() -> [value:this is a test string] [status:OK]
	//test: Resolve() -> [value:this is a test string] [status:OK]

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
	fragment := "v2"

	status := Resolver.AddRepresentation(name, "author", messaging.Content{Fragment: fragment, Type: messaging.ContentTypeJson, Value: addr})
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name, fragment)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
	}

	s3, status3 := Resolve[Address](name, fragment, Resolver)
	fmt.Printf("test: Resolve() -> [value:%v] [status:%v]\n", s3, status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment: v2 type: application/json value: true] [status:OK]
	//test: Representation() -> [value:77] [status:OK]
	//test: Resolve() -> [value:{123 Main  Anytown Ohio 54321}] [status:OK]

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
	fragment := "v2"

	status := Resolver.AddRepresentation(name, "author", messaging.Content{Fragment: fragment, Type: messaging.ContentTypeJson, Value: m})
	fmt.Printf("test: AddRepresentation() -> [status:%v]\n", status)

	ct, status2 := Resolver.Representation(name, fragment)
	fmt.Printf("test: Representation() -> [ct:%v] [status:%v]\n", ct, status2)

	if buf, ok := ct.Value.([]byte); ok {
		fmt.Printf("test: Representation() -> [value:%v] [status:%v]\n", len(buf), status2)
	}

	s3, status3 := Resolve[map[string]string](name, fragment, Resolver)
	fmt.Printf("test: Resolve() -> [value:%v] [status:%v]\n", s3, status3)

	//Output:
	//test: AddRepresentation() -> [status:OK]
	//test: Representation() -> [ct:fragment: v2 type: application/json value: true] [status:OK]
	//test: Representation() -> [value:77] [status:OK]
	//test: Resolve() -> [value:map[City:Anytown Line1:123 Main Line2: State:Ohio Zip:54321]] [status:OK]

}
