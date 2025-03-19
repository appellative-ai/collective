package namespace

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func httpPutThing(nsName, author string) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}

func httpPutRelation(nsName1, nsName2, author string) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}
