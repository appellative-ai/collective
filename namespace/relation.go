package namespace

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/core/std"
	"net/http"
)

type tagRelation struct {
	Name     string `json:"name"`
	Instance string `json:"instance"`
	Pattern  string `json:"pattern"`
	Args     []Arg  `json:"args"`
}

// TODO: additional validation for args
func (a *agentT) relation(name string, args []Arg) *std.Status {
	if name == "" {
		return std.NewStatus(http.StatusBadRequest, "", errors.New(fmt.Sprintf("empty name")))
	}
	//status := a.intf.Thing(http.MethodPut, name, cname, author)
	//if !status.OK() {
	//	return status //.WithMessage(fmt.Sprintf("name %v", name))
	//}
	return std.StatusOK
}
