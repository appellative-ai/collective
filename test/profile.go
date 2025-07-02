package test

import (
	"github.com/behavioral-ai/collective/resource"
	"github.com/behavioral-ai/core/messaging"
)

const (
	ProfileName = "behavioral-ai:resiliency:type/domain/metrics/profile"
)

func LoadProfile(r *resource.Resolution) *messaging.Status {
	//url, _ := url.Parse(testfs.ResiliencyTrafficProfile1)
	return r.AddRepresentation(ProfileName, "", "author", "") //resource.Content{})
}
