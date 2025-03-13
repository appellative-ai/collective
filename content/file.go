package content

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

// fileResolution - is read only and returns "not found" on gets
func fileResolution(method, name, _ string, _ []byte, version int) ([]byte, *messaging.Status) {
	// file resolution is read only
	if method == http.MethodPut {
		return nil, messaging.StatusOK()
	}
	return nil, messaging.StatusNotFound()
}
