package namespace

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/iox"
	"github.com/appellative-ai/core/std"
	"net/http"
	"time"
)

type retrieval struct {
	Name string `json:"name"`
	Args []Arg  `json:"args"`
}

func (a *agentT) retrieval(name string, args []Arg) (*std.Content, *std.Status) {
	c := std.Content{}

	buf, err := createRetrieval(name, args)
	if err != nil {
		return nil, std.NewStatus(http.StatusBadRequest, a.Name(), err)
	}
	newCtx, cancel := httpx.NewContext(nil, a.timeout)
	defer cancel()
	req, err1 := http.NewRequestWithContext(newCtx, http.MethodPost, a.url(retrievalPath), bytes.NewBuffer(buf))
	if err1 != nil {
		return nil, std.NewStatus(http.StatusInternalServerError, a.Name(), err1)
	}
	start := time.Now().UTC()
	resp, err2 := a.ex(req)
	a.log(start, time.Since(start), retrievalRoute, req, resp, a.timeout)
	if err2 != nil {
		return nil, std.NewStatus(http.StatusInternalServerError, a.Name(), err2)
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, std.NewStatus(resp.StatusCode, a.Name(), nil)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, std.NewStatus(resp.StatusCode, a.Name(), nil)
	}
	ok, err3 := createContent(resp, &c)
	if err3 != nil {
		return nil, std.NewStatus(http.StatusInternalServerError, a.Name(), err3)
	}
	if !ok {
		return nil, std.StatusNotFound
	}
	return &c, std.NewStatus(resp.StatusCode, a.Name(), nil)
}

// TODO: add validation for required arguments
func createRetrieval(name string, args []Arg) ([]byte, error) {
	if name == "" {
		return nil, errors.New(fmt.Sprintf("empty name "))
	}
	if len(args) == 0 {
		return nil, errors.New(fmt.Sprintf("empty args"))
	}
	return json.Marshal(&retrieval{Name: name, Args: args})
}

func createContent(resp *http.Response, content *std.Content) (bool, error) {
	buf, err := iox.ReadAll(resp.Body, resp.Header)
	if err != nil {
		return false, err
	}
	if len(buf) == 0 {
		return false, nil
	}
	content.Type = resp.Header.Get("Content-Type")
	content.Value = buf
	return true, nil
}
