package content

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

var (
	NotFound = NewError(http.StatusNotFound, nil, "")
)

type Error struct {
	Code int    `json:"code"`
	Err  error  `json:"err"`
	Msg  string `json:"message"`
}

func NewError(code int, err error, msg string) *Error {
	e := new(Error)
	e.Code = code
	e.Err = err
	e.Msg = msg
	return e
}

func (e *Error) String() string {
	status := messaging.HttpStatus(e.Code)
	if e.Msg != "" {
		return fmt.Sprintf("%v [msg:%v]", status, e.Msg)
	}
	if e.Err == nil {
		return fmt.Sprintf("%v", status)
	}
	return fmt.Sprintf("%v [err:%v]", status, e.Err)
}

func (e *Error) SetMessage(msg string) *Error {
	e.Msg = msg
	return e
}
