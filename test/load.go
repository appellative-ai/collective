package test

import (
	"github.com/behavioral-ai/collective/resource"
	"github.com/behavioral-ai/collective/testfs"
	"github.com/behavioral-ai/core/messaging"
	"net/url"
)

func loadResolver(resolver *resource.Resolution) *messaging.Status {
	url, _ := url.Parse(testfs.ResiliencyInterpret1)
	status := resolver.AddRepresentation(ResiliencyInterpret, "", "author", "") //resource.Content{})
	if !status.OK() {
		return status
	}
	url, _ = url.Parse(testfs.ResiliencyThreshold1)
	return resolver.AddRepresentation(ResiliencyThreshold, "", "author", "") //resource.Content{})
}
