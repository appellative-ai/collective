package content

import (
	"context"
	"github.com/behavioral-ai/core/io"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"net/url"
	"strconv"
)

func httpGetContent(nsName string, version int) ([]byte, *messaging.Status) {
	v := make(url.Values)
	v.Set(NsNameKey, nsName)
	v.Set(VersionKey, strconv.Itoa(version))
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://collective/content?"+v.Encode(), nil)
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
	return buf, nil
}

func httpPutContent(nsName, author string, value []byte, version int) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}

func httpAddActivity(hostName, uri, event, source string, content any) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}

func httpNotify(e messaging.Event) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}
