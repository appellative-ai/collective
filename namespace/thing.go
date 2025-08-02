package namespace

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/std"
	"net/http"
	"time"
)

type thing struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Author string `json:"author"`
}

func (a *agentT) addThing(name, cname, author string) *std.Status {
	buf, err := createThing(name, cname, author)
	if err != nil {
		return std.NewStatus(http.StatusBadRequest, a.Name(), err)
	}
	newCtx, cancel := httpx.NewContext(nil, a.timeout)
	defer cancel()
	req, err1 := http.NewRequestWithContext(newCtx, http.MethodPost, a.url(requestThingPath), bytes.NewBuffer(buf))
	if err1 != nil {
		return std.NewStatus(http.StatusInternalServerError, a.Name(), err1)
	}
	start := time.Now().UTC()
	resp, err2 := a.ex(req)
	a.log(start, time.Since(start), requestThingRoute, req, resp, a.timeout)
	if err2 != nil {
		return std.NewStatus(http.StatusInternalServerError, a.Name(), err2)
	}
	return std.NewStatus(resp.StatusCode, a.Name(), nil)
}

func createThing(name, cname, author string) ([]byte, error) {
	if name == "" {
		return nil, errors.New(fmt.Sprintf("empty name"))
	}
	if author == "" {
		return nil, errors.New(fmt.Sprintf("empty author"))
	}

	return json.Marshal(&link{Name: name, CName: cname, Author: author})
}
