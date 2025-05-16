package test

import (
	"github.com/behavioral-ai/collective/fs"
	"github.com/behavioral-ai/core/messaging"
	"net/url"
)

func loadResolver(resolver *resource.Resolution) *messaging.Status {
	url, _ := url.Parse(fs.ResiliencyInterpret1)
	status := resolver.Add(ResiliencyInterpret, "author", url)
	if !status.OK() {
		return status
	}
	url, _ = url.Parse(fs.ResiliencyThreshold1)
	return resolver.Add(ResiliencyThreshold, "author", url)
}
