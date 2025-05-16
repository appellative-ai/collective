package test

import (
	"github.com/behavioral-ai/collective/fs"
	"github.com/behavioral-ai/core/messaging"
	"net/url"
)

const (
	ProfileName = "behavioral-ai:resiliency:type/domain/metrics/profile"
)

func LoadProfile(r *resource.Resolution) *messaging.Status {
	url, _ := url.Parse(fs.ResiliencyTrafficProfile1)
	return r.Add(ProfileName, "author", url)
}
