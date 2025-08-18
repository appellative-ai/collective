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

func Startup(cfg map[string]string,
	status func(status any),
	exchange func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) error {
	return nil
}
