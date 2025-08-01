package namespace

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/core/std"
	"net/http"
)

type tagThing struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Author string `json:"author"`
}

func (a *agentT) addThing(name, cname, author string) *std.Status {
	if name == "" || author == "" {
		return std.NewStatus(http.StatusBadRequest, "", errors.New(fmt.Sprintf("empty name [%v] or author [%v]", name, author)))
	}
	//status := a.intf.Thing(http.MethodPut, name, cname, author)
	//if !status.OK() {
	//	return status //.WithMessage(fmt.Sprintf("name %v", name))
	//}
	return std.StatusOK
}
