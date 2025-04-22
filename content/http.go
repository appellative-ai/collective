package content

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// TODO : support HEAD requests so that variants, of different content type, can be supported
func httpGetContent(nsName, resource, version string) (Accessor, *messaging.Status) {
	/*
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, module.ContentURL(nsName, version), nil)
		if err != nil {
			return nil, messaging.NewStatusError(messaging.StatusInvalidArgument, err, AgentNamespaceName)
		}
		resp, err1 := Exchange(req)
		if err1 != nil {
			return nil, messaging.NewStatusError(resp.StatusCode, err1, AgentNamespaceName)
		}
		if resp.StatusCode != http.StatusOK {
			return nil, messaging.NewStatus(resp.StatusCode)
		}
		buf, err2 := io.ReadAll(resp.Body, resp.Header)
		if err2 != nil {
			return nil, messaging.NewStatusError(resp.StatusCode, err2, AgentNamespaceName)
		}
	*/
	return Accessor{}, messaging.StatusNotFound()
}

func httpPutContent(nsName, resource, version, author string, access Accessor) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}
