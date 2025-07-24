package test

import (
	"github.com/appellative-ai/collective/resolution"
	"github.com/appellative-ai/collective/testfs"
	"github.com/appellative-ai/core/messaging"
	"net/url"
)

func loadResolver(resolver *resolution.Resolution) *messaging.Status {
	url, _ := url.Parse(testfs.ResiliencyInterpret1)
	status := resolver.AddRepresentation(ResiliencyInterpret, "", "author", "") //resolution.Content{})
	if !status.OK() {
		return status
	}
	url, _ = url.Parse(testfs.ResiliencyThreshold1)
	return resolver.AddRepresentation(ResiliencyThreshold, "", "author", "") //resolution.Content{})
}
