package operations

import (
	"errors"
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

func Origin() *OriginT {
	return agent.origin
}

func ConfigOrigin(cfg map[string]string) error {
	if len(cfg) == 0 {
		return errors.New("empty origin")
	}
	var err error
	var o OriginT
	err = newOrigin(&o, cfg)
	if err != nil {
		return err
	}
	agent.origin = &o
	return nil
}

func ConfigRegistryHosts(hosts []string) error {
	if len(hosts) == 0 || hosts[0] == "" {
		return errors.New("registry hosts are required")
	}
	s := agent.state.Load()
	s.registryHosts = hosts
	return nil
}

func ConfigLogging(logFunc func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) {
	if logFunc != nil {
		agent.logFunc = logFunc
	}
}

func Startup() error {
	return agent.startup()
}
