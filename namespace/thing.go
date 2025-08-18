package namespace

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/appellative-ai/core/std"
	"net/http"
)

type thing struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Author string `json:"author"`
}

func (a *agentT) addThing(name, cname, author string) *std.Status {
	buf, err := createThing(name, cname, author)
	if err != nil {
		return std.NewStatusWithLocation(http.StatusBadRequest, err, a.Name())
	}
	resp, status := a.call(http.MethodPost, a.url(requestThingPath), requestThingRoute, nil, buf)
	if !status.OK() {
		return status
	}
	return std.NewStatus(resp.StatusCode, nil)
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
