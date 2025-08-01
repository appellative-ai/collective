package namespace

import (
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
	if name == "" || thing1 == "" || thing2 == "" || author == "" {
		return std.NewStatus(http.StatusBadRequest, "", errors.New(fmt.Sprintf("empty name [%v], thing1 [%v], thing2 [%v], author [%v]", name, thing1, thing2, author)))
	}
	// TODO: remove after initial testing
	/*
		a.relations.put(name, thing1, thing2)

		status := a.intf.Relation(http.MethodPut, name, cname, thing1, thing2, author)
		if !status.OK() {
			return status //.WithMessage(fmt.Sprintf("name1 %v", name))
		}

	*/
	return std.StatusOK
}
