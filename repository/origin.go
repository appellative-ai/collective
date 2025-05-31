package repository

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

const (
	NamespaceName = "core:type/repository/origin"
	RegionKey     = "region"
	ZoneKey       = "zone"
	SubZoneKey    = "sub-zone"
	HostKey       = "host"
	InstanceIdKey = "instance-id"
	uriFmt        = "%v:%v"
)

// Origin - location
type Origin struct {
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`
}

func (o Origin) Uri(class string) string {
	return fmt.Sprintf(uriFmt, class, o)
}

func (o Origin) String() string {
	var uri = o.Region

	if o.Zone != "" {
		uri += "." + o.Zone
	}
	if o.SubZone != "" {
		uri += "." + o.SubZone
	}
	if o.Host != "" {
		uri += "." + o.Host
	}
	return uri
}

func NewOriginFromMessage(m *messaging.Message) (o Origin, ok bool) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyMapError(NamespaceName), NamespaceName)
		return
	}
	o.Region = cfg[RegionKey]
	if o.Region == "" {
		messaging.Reply(m, messaging.ConfigMapContentError(NamespaceName, RegionKey), NamespaceName)
		return
	}

	o.Zone = cfg[ZoneKey]
	if o.Zone == "" {
		messaging.Reply(m, messaging.ConfigMapContentError(NamespaceName, ZoneKey), NamespaceName)
		return
	}

	o.SubZone = cfg[SubZoneKey]
	//if o.SubZone == "" {
	//	messaging.Reply(m, messaging.ConfigMapContentError(nil, SubZoneKey), NamespaceName)
	//	return
	//}

	o.Host = cfg[HostKey]
	if o.Host == "" {
		messaging.Reply(m, messaging.ConfigMapContentError(NamespaceName, HostKey), NamespaceName)
		return
	}
	o.InstanceId = cfg[InstanceIdKey]
	/*
		if o.InstanceId == "" {
			messaging.Reply(m, messaging.ConfigMapContentError(a, InstanceIdKey), a.Name())
			return
		}
	*/
	messaging.Reply(m, messaging.StatusOK(), NamespaceName)
	return o, true
}
