package http

import (
	"fmt"
	"github.com/behavioral-ai/collective/testfs"
	"github.com/behavioral-ai/core/httpx"
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
	var name string

	switch r.URL.Path {
	case textResource:
		name = testfs.BehavioralAITextExample
	case htmlResource:
		name = testfs.BehavioralAIHtmlExample
	case jsonResource:
		name = testfs.BehavioralAIJsonExample
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	buf, err := testfs.ReadFile(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add(httpx.ContentType, http.DetectContentType(buf))
	w.Header().Add(httpx.ContentLength, fmt.Sprintf("%v", len(buf)))
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
