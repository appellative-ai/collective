package test

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/testrsc"
	"github.com/behavioral-ai/core/messaging"
	"net/url"
)

func loadResolver(resolver *content.Resolution) *messaging.Status {
	url, _ := url.Parse(testrsc.ResiliencyInterpret1)
	status := resolver.Add(ResiliencyInterpret, "author", url)
	if !status.OK() {
		return status
	}
	url, _ = url.Parse(testrsc.ResiliencyThreshold1)
	return resolver.Add(ResiliencyThreshold, "author", url)
}
