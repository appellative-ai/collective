package operations

import (
	"github.com/appellative-ai/core/std"
)

func thing(method, name, cname, author string) *std.Status {
	return std.StatusOK
}

func relation(method, name, cname, thing1, thing2, author string) *std.Status {
	return std.StatusOK
}
