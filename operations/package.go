package operations

import (
	"errors"
	"github.com/appellative-ai/core/messaging"
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

func Startup() error {
	return agent.startup()
}

func Shutdown() {
	agent.Message(messaging.ShutdownMessage)
}
