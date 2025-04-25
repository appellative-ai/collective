package test

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/testrsc"
	"github.com/behavioral-ai/core/messaging"
	"net/url"
)

const (
	ProfileName = "behavioral-ai:resiliency:type/domain/metrics/profile"
)

func LoadProfile(r *content.Resolution) *messaging.Status {
	url, _ := url.Parse(testrsc.ResiliencyTrafficProfile1)
	return r.Add(ProfileName, "author", url)
}
