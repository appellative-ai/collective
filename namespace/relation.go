package namespace

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/appellative-ai/core/std"
	"net/http"
)

type relation struct {
	Name string `json:"name"`
	Args []Arg  `json:"args"`
}

func (a *agentT) relation(name string, args []Arg) (*std.Content, *std.Status) {
	buf, err := createRelation(name, args)
	if err != nil {
		return nil, std.NewStatus(http.StatusBadRequest, a.Name(), err)
	}
	resp, status := a.call(http.MethodPost, a.url(relationPath), relationRoute, nil, buf)
	if !status.OK() {
		return nil, status
	}
	content, err3 := createContent(resp)
	if err3 != nil {
		return nil, std.NewStatus(http.StatusInternalServerError, a.Name(), err3)
	}
	return content, std.NewStatus(resp.StatusCode, a.Name(), nil)
}

// TODO: add validation for required arguments
func createRelation(name string, args []Arg) ([]byte, error) {
	if name == "" {
		return nil, errors.New(fmt.Sprintf("empty name "))
	}
	if len(args) == 0 {
		return nil, errors.New(fmt.Sprintf("empty args"))
	}
	return json.Marshal(&relation{Name: name, Args: args})
}
