package operations

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

const (
	uriFmt = "%v:%v"
)

// origin - location
type originT struct {
	AppId      string `json:"app-id"`
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`
}

func (o originT) Uri(class string) string {
	return fmt.Sprintf(uriFmt, class, o)
}

func (o originT) String() string { return o.Name() }

/*var uri = o.Region

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

*/

func (o originT) Name() string {
	var name = o.AppId + ":" + ServiceKind

	if o.Region != "" {
		name += "/" + o.Region
	}
	if o.Zone != "" {
		name += "/" + o.Zone
	}
	if o.SubZone != "" {
		name += "/" + o.SubZone
	}
	if o.Host != "" {
		name += "/" + o.Host
	}
	if o.InstanceId != "" {
		name += "#" + o.InstanceId
	}
	return name
}

func newOriginFromMessage(m *messaging.Message) (o originT, ok bool) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyMapError(NamespaceName), NamespaceName)
		return
	}
	o.AppId = cfg[AppIdKey]
	if o.AppId == "" {
		messaging.Reply(m, messaging.ConfigMapContentError(NamespaceName, AppIdKey), NamespaceName)
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
