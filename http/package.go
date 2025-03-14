package http

import "net/http"

// Exchange - exchange type
type Exchange func(r *http.Request) (*http.Response, error)
