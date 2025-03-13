package test

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/testrsc"
	"github.com/behavioral-ai/core/messaging"
	url2 "net/url"
)

const (
	ProfileName = "resiliency:type/domain/metrics/profile"
)

func LoadProfile(r content.Resolution) *messaging.Status {
	url, _ := url2.Parse(testrsc.ResiliencyTrafficProfile1)
	return r.PutValue(ProfileName, "author", url, 1)
}
