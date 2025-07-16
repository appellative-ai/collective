package http

import (
	"fmt"
	"github.com/appellative-ai/collective/testfs"
	"net/http"
	"strings"
)

const (
	CollectivePattern = "/collective"
	textResource      = "/collective:appellative-ai:text"
	htmlResource      = "/collective:appellative-ai:html"
	jsonResource      = "/collective:appellative-ai:json"
	contentType       = "Content-Type"
	contentLength     = "Content-Length"
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
	w.Header().Add(contentType, http.DetectContentType(buf))
	w.Header().Add(contentLength, fmt.Sprintf("%v", len(buf)))
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
