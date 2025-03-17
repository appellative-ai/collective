package module

import (
	"net/url"
	"strconv"
)

const (
	ContentPath    = "/collective/content"
	ActivityPath   = "/collective/activity"
	NotifyPath     = "/collective/notify"
	NamespacePath  = "/collective/namespace"
	TimeseriesPath = "/collective/timeseries"
	NsNameKey      = "name"
	VersionKey     = "ver"
)

var (
	HostName string
)

func Initialize(hostName string) {
	HostName = hostName
}

func ContentURL(nsName string, version int) string {
	v := make(url.Values)
	v.Set(NsNameKey, nsName)
	v.Set(VersionKey, strconv.Itoa(version))
	return ContentPath + "?" + v.Encode()
}

func ActivityURL() string {
	return ActivityPath
}

func NotifyURL() string {
	return NotifyPath
}
