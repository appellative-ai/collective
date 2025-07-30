package namespace

import (
	"encoding/json"
	"fmt"
)

func ExampleRelation() {
	r := tagRelation{
		Name:     "test:agent",
		Instance: "core:aspect/resiliency",
		Pattern:  "core:aspect/expressive",
		Args: []Arg{
			{Name: "kind", Value: "aspect"},
			{Name: "count", Value: 25},
		},
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Relation() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Relation() -> [{"name":"test:agent","instance":"core:aspect/resiliency","pattern":"core:aspect/expressive","args":[{"name":"kind","value":"aspect"},{"name":"count","value":25}]}] [err:<nil>]

}

func ExampleRetrieval() {
	r := tagRetrieval{
		Name: "test:agent",

		Args: []Arg{
			{Name: "kind", Value: "aspect"},
			{Name: "count", Value: 25},
		},
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Retrieval() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Retrieval() -> [{"name":"test:agent","args":[{"name":"kind","value":"aspect"},{"name":"count","value":25}]}] [err:<nil>]

}

func ExampleThing() {
	r := tagThing{
		Name:   "test:agent",
		CName:  "headers",
		Author: "core:person/bob",
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Thing() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Thing() -> [{"name":"test:agent","cname":"headers","author":"core:person/bob"}] [err:<nil>]

}

func ExampleLink() {
	r := tagLink{
		Name:   "test:agent",
		CName:  "headers",
		Author: "core:person/bob",
		Thing1: "core:thing/first",
		Thing2: "core:thing/second",
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Link() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Link() -> [{"name":"test:agent","cname":"headers","thing1":"core:thing/first","thing2":"core:thing/second","author":"core:person/bob"}] [err:<nil>]
	
}
