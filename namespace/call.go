package namespace

import (
	"bytes"
	"errors"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/iox"
	"github.com/appellative-ai/core/std"
	"io"
	"net/http"
	"time"
)

func (a *agentT) call(method, uri, route string, h http.Header, body []byte) (*http.Response, *std.Status) {
	ctx, cancel := httpx.NewContext(nil, a.timeout)
	defer cancel()

	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewBuffer(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, uri, reqBody)
	if err != nil {
		return httpx.NewResponse(http.StatusBadRequest, nil, nil), std.NewStatus(http.StatusBadRequest, a.Name(), err)
	}
	start := time.Now().UTC()
	req.Header = h
	resp, err2 := a.exchange(req)
	a.log(start, time.Since(start), route, req, resp, a.timeout)
	if err2 != nil {
		return resp, std.NewStatus(http.StatusInternalServerError, a.Name(), err2)
	}
	return resp, std.NewStatus(resp.StatusCode, a.Name(), nil)
}

func createContent(resp *http.Response) (*std.Content, error) {
	buf, err := iox.ReadAll(resp.Body, resp.Header)
	if err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		return nil, nil
	}
	if resp.Header == nil {
		return nil, errors.New("nil header")
	}
	return &std.Content{
		Fragment: "",
		Type:     resp.Header.Get(httpx.ContentType),
		Value:    buf,
	}, nil
}
