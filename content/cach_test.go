package content

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ExampleNewCache() {
	name := "test:thing:text"

	c := newCache()
	buf, err := json.Marshal(text{Value: "Hello World!"})
	if err != nil {
		fmt.Printf("test: json.Marshall() -> [err:%v]\n", err)
	} else {
		//var status error //status *messaging.Status
		//	var access Accessor{}
		c.put(name, "", Accessor{Version: "", Type: http.DetectContentType(buf), Content: buf})
		//fmt.Printf("test: newContentCache.put(1) -> [status:%v]\n", status)

		access, status := c.get(name, "")
		fmt.Printf("test: newContentCache.get(\"%v\") -> [%v] [status:%v]\n", name, access, status)

		//name = name + "#1.2.4"
		_, status = c.get(name, "#1.2.4")
		fmt.Printf("test: newContentCache.get(\"%v\") -> [%v] [status:%v]\n", name, access, status)

		var v text
		err = json.Unmarshal(buf, &v)
		fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, v)
	}

	//Output:
	//test: newContentCache.get("test:thing:text") -> [vers:  type: text/plain; charset=utf-8 content: true] [status:<nil>]
	//test: newContentCache.get("test:thing:text") -> [vers:  type: text/plain; charset=utf-8 content: true] [status:content [test:thing:text] not found]
	//test: json.Unmarshal() -> [err:<nil>] [{Hello World!}]

}
