package http

import (
	"fmt"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/iox"
	"net/http"
	"net/http/httptest"
)

func ExampleHandler() {
	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8081"+textResource, nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)
	rec.Flush()
	fmt.Printf("test: Handler() -> [status:%v] [ct:%v] [len:%v]\n", rec.Result().StatusCode, rec.Result().Header.Get(httpx.ContentType), rec.Result().Header.Get(httpx.ContentLength))

	buf, err := iox.ReadAll(rec.Result().Body, nil)
	fmt.Printf("test: Handler() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Handler() -> [status:200] [ct:text/plain; charset=utf-8] [len:1649]

}
