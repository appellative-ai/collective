package namespace

import "fmt"

func ExampleParse() {
	s := "wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http#v1.2.3"
	name := parse(s)
	fmt.Printf("test: parse(\"%v\") -> %v\n", s, name)

	s = "wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http"
	name = parse(s)
	fmt.Printf("test: parse(\"%v\") -> %v\n", s, name)

	//Output:
	//test: parse("wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http#v1.2.3") -> {wikipedia.eng resiliency.traffic agent /rate-limiting/request/http #v1.2.3}
	//test: parse("wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http") -> {wikipedia.eng resiliency.traffic agent /rate-limiting/request/http }

}

func ExampleAddVersion() {
	s := "wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http#v1.2.3"
	name := addVersion(s, "v3,4,5")
	fmt.Printf("test: addVersion(\"%v\") -> %v\n", s, name)

	s = "wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http"
	name = addVersion(s, "v9.8.7")
	fmt.Printf("test: addVersion(\"%v\") -> %v\n", s, name)

	//Output:
	//test: addVersion("wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http#v1.2.3") -> wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http#v1.2.3
	//test: addVersion("wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http") -> wikipedia.eng:resiliency.traffic:agent/rate-limiting/request/http#v9.8.7

}
