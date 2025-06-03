package namespace

import (
	"net/url"
	"strings"
)

func parse(name string) Name {
	if name == "" {
		return Name{}
	}
	u, err := url.Parse(name)
	if err != nil {
		return Name{Collective: err.Error()}
	}
	n := Name{Collective: u.Scheme, Fragment: u.Fragment}
	i := strings.Index(u.Opaque, Colon)
	if i < 0 {
		n.Collective = "error, missing second colon"
		return n
	}
	n.Domain = u.Opaque[:i]
	path := u.Opaque[i:]
	i = strings.Index(path, Slash)
	if i < 0 {
		return n
	}
	n.Kind = path[1:i]
	n.Path = path[i:]
	return n
}

func addVersion(name, version string) string {
	if name == "" {
		return name
	}
	i := strings.Index(name, Fragment)
	if i == -1 {
		return name + Fragment + version
	}
	return name
}
