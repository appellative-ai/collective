package operations

import (
	"time"
)

// Origin map and host keys
const (
	RegionKey        = "region"
	ZoneKey          = "zone"
	SubZoneKey       = "sub-zone"
	HostKey          = "host"
	InstanceIdKey    = "instance-id"
	ServiceNameKey   = "service-name"
	CollectiveKey    = "collective"
	DomainKey        = "domain"
	RegistryHost1Key = "registry-host1"
	RegistryHost2Key = "registry-host2"
)

var (
	Origin OriginT
)

func Startup(cfg map[string]string, log func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) error {
	var err error
	Origin, err = NewOrigin(cfg)
	if err != nil {
		return err
	}
	return agent.startup(cfg[CollectiveKey], cfg[RegistryHost1Key], cfg[RegistryHost2Key], log)
}
