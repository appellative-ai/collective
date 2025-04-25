package http

import (
	"fmt"
	"github.com/behavioral-ai/collective/fs"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/iox"
	"net/http"
	"strings"
)

const (
	CollectivePattern = "/collective"
	textResource      = "/collective:behavioral-ai:text"
	htmlResource      = "/collective:behavioral-ai:html"
	jsonResource      = "/collective:behavioral-ai:json"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, CollectivePattern) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var file string

	switch r.URL.Path {
	case textResource:
		file = fs.BehavioralAITextExample
	case htmlResource:
		file = fs.BehavioralAIHtmlExample
	case jsonResource:
		file = fs.BehavioralAIJsonExample
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	buf, err := iox.ReadFile(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add(httpx.ContentType, http.DetectContentType(buf))
	w.Header().Add(httpx.ContentLength, fmt.Sprintf("%v", len(buf)))
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
