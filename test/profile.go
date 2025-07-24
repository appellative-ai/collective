package test

import (
	"github.com/appellative-ai/collective/resolution"
	"github.com/appellative-ai/core/std"
)

const (
	ProfileName = "appellative-ai:resiliency:type/domain/metrics/profile"
)

func LoadProfile(r *resolution.Interface) *std.Status {
	/*
		    //url, _ := url.Parse(testfs.ResiliencyTrafficProfile1)
			return r.AddRepresentation(ProfileName, "", "author", "") //resolution.Content{})

	*/
	return std.StatusOK
}
