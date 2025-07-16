package operations

import "github.com/appellative-ai/core/messaging"

func thing(method, name, cname, author string) *messaging.Status {
	return messaging.StatusOK()
}

func relation(method, name, cname, thing1, thing2, author string) *messaging.Status {
	return messaging.StatusOK()
}
