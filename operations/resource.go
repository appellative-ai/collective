package operations

import (
	"github.com/appellative-ai/core/std"
)

func representation(method, name, author, contentType string, value []byte) (std.Content, *std.Status) {
	return std.Content{}, std.StatusOK
}

func context(method, name, author, contentType string, value []byte) (std.Content, *std.Status) {
	return std.Content{}, std.StatusOK
}
