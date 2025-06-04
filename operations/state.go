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
	origin            messaging.Origin
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
	var err error
	if ops.origin, err = messaging.NewOriginFromMessage(msg); err != nil {
		// TODO: reply with error
		return
	}
	ops.serviceName = ops.collective + ":" + ops.origin.Name(ops.collective, ops.domain)
	return
}
