package content

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

var (
	NotFound = NewStatus(http.StatusNotFound, nil)
)

type Status struct {
	Code int    `json:"code"`
	Err  error  `json:"err"`
	Msg  string `json:"message"`
}

func StatusOK() *Status {
	return okStatus
}

func NewStatus(code int, msg any) *Status {
	s := new(Status)
	s.Code = code
	if msg == nil {
		return s
	}
	if e, ok := msg.(error); ok {
		s.Err = e
		return s
	}
	if str, ok := msg.(string); ok {
		s.Msg = str
	}
	return s
}

func (s *Status) OK() bool {
	return s.Code == http.StatusOK
}

func (s *Status) NotFound() bool {
	return s.Code == http.StatusNotFound
}

func (s *Status) NoContent() bool {
	return s.Code == http.StatusNoContent
}

func (s *Status) String() string {
	status := messaging.HttpStatus(s.Code)
	if s.Msg != "" {
		return fmt.Sprintf("%v [msg:%v]", status, s.Msg)
	}
	if s.Err == nil {
		return fmt.Sprintf("%v", status)
	}
	return fmt.Sprintf("%v [err:%v]", status, s.Err)
}

func (s *Status) SetMessage(msg string) *Status {
	s.Msg = msg
	return s
}

var okStatus = func() *Status {
	s := new(Status)
	s.Code = http.StatusOK
	return s
}()
