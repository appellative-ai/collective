package content

import (
	"errors"
	"fmt"
	"sync"
)

// resolutionKey -
type resolutionKey struct {
	Name     string `json:"name"`
	Resource string `json:"resource"`
	Version  string `json:"version"`
}

type content struct {
	body Accessor
}

type contentT struct {
	m *sync.Map
}

func newContentCache() *contentT {
	c := new(contentT)
	c.m = new(sync.Map)
	return c
}

func (c *contentT) get(name, resource, version string) (Accessor, error) {
	key := resolutionKey{Name: name, Resource: resource, Version: version}
	value, ok := c.m.Load(key)
	if !ok {
		return Accessor{}, errors.New(fmt.Sprintf("content [%v] [%v] not found", name, version))
	}
	if value1, ok1 := value.(content); ok1 {
		return value1.body, nil
	}
	return Accessor{}, nil
}

func (c *contentT) put(name, resource, version string, access Accessor) {
	c.m.Store(resolutionKey{Name: name, Resource: resource, Version: version}, content{body: access})
}
