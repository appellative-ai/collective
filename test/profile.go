package test

import (
	"github.com/appellative-ai/collective/resource"
	"github.com/appellative-ai/core/messaging"
)

const (
	ProfileName = "appellative-ai:resiliency:type/domain/metrics/profile"
)

func LoadProfile(r *resource.Resolution) *messaging.Status {
	//url, _ := url.Parse(testfs.ResiliencyTrafficProfile1)
	return r.AddRepresentation(ProfileName, "", "author", "") //resource.Content{})
}
