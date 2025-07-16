package test

import (
	"github.com/appellative-ai/collective/resource"
	"github.com/appellative-ai/collective/testfs"
	"github.com/appellative-ai/core/messaging"
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
