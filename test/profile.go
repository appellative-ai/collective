package test

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/testrsc"
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
)

const (
	ProfileName = "resiliency:type/domain/metrics/profile"
)

func LoadProfile(r content.Resolution) *messaging.Status {
	buf, err := iox.ReadFile(testrsc.ResiliencyTrafficProfile1)
	if err != nil {
		return messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	return r.PutContent(ProfileName, "author", buf, 1)
}
