package namespace

import "strings"

func parse(name string) Name {
	if name == "" {
		return Name{}
	}
	tokens := strings.Split(name, Colon)
	if len(tokens) < 3 {
		return Name{Collective: "error invalid name"}
	}
	n := Name{Collective: tokens[0], Domain: tokens[1]}
	i := strings.Index(tokens[2], Slash)
	if i < 0 {
		return n
	}
	n.Kind = tokens[2][:i]
	path := tokens[2][i:]
	i = strings.Index(path, Fragment)
	if i == -1 {
		n.Path = path
	} else {
		n.Path = path[:i]
		n.Fragment = path[i:]
	}
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
