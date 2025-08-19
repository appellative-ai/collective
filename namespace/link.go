package namespace

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/appellative-ai/core/std"
	"net/http"
)

type link struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
	Author string `json:"author"`
}

func (a *agentT) addLink(name, cname, thing1, thing2, author string) *std.Status {
	buf, err := createLink(name, cname, thing1, thing2, author)
	if err != nil {
		return std.NewStatus(http.StatusBadRequest, err).SetLocation(a.Name())
	}
	resp, status := a.call(http.MethodPost, a.url(requestLinkPath), requestLinkRoute, nil, buf)
	if !status.OK() {
		return status
	}
	return std.NewStatus(resp.StatusCode, nil)
}

func createLink(name, cname, thing1, thing2, author string) ([]byte, error) {
	if name == "" || author == "" {
		return nil, errors.New(fmt.Sprintf("empty name [%v] or author [%v]", name, author))
	}
	if thing1 == "" || thing2 == "" {
		return nil, errors.New(fmt.Sprintf("empty thing1 [%v] or empty thing2 [%v]", thing1, thing2))
	}
	return json.Marshal(&link{Name: name, CName: cname, Thing1: thing1, Thing2: thing2, Author: author})
}
