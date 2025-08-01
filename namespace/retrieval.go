package namespace

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/core/std"
	"net/http"
)

type tagRetrieval struct {
	Name string `json:"name"`
	Args []Arg  `json:"args"`
}

// TODO: additional validation for args
func (a *agentT) retrieval(name string, args []Arg) (*std.Content, *std.Status) {
	if name == "" {
		return nil, std.NewStatus(http.StatusBadRequest, "", errors.New(fmt.Sprintf("empty name")))
	}
	//status := a.intf.Thing(http.MethodPut, name, cname, author)
	//if !status.OK() {
	//	return status //.WithMessage(fmt.Sprintf("name %v", name))
	//}
	return nil, std.StatusOK
}
