package namespace

import (
	"net/http"
)

func httpPutThing(name, cname, authority, author string) (*http.Response, *Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, StatusOK()
}

func httpPutRelation(name, cname, thing1, thing2, authority, author string) (*http.Response, *Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, StatusOK()
}
