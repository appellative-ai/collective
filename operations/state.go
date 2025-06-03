package operations

import "github.com/behavioral-ai/core/messaging"

type operationsT struct {
	running           bool
	registryHost1     string // registry host name
	registryHost2     string // registry host name
	collective        string
	domain            string
	serviceName       string
	linkedCollectives map[string][]string
	origin            originT
}

// TODO: need to resolve all of the links in a collective and query the registry for the
//
//	host names for the collective
//
//	Need a default domain for metadata/links -> root??, import??
func initialize(msg *messaging.Message) (ops *operationsT, ok bool) {
	cfg := messaging.ConfigMapContent(msg)
	ops = new(operationsT)
	ops.registryHost1 = cfg[RegistryHost1Key]
	ops.registryHost2 = cfg[RegistryHost2Key]
	ops.collective = cfg[CollectiveKey]
	ops.domain = cfg[DomainKey]
	if ops.origin, ok = newOriginFromMessage(msg); !ok {
		return
	}
	ops.serviceName = ops.collective + ":" + ops.origin.Name()
	return
}
