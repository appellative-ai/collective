package resource

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"sync"
)

// resolutionKey -
type resolutionKey struct {
	Name     string `json:"name"`
	Fragment string `json:"fragment"`
}

/*
type content struct {
	body messaging.Content
}


*/

type cacheT struct {
	m *sync.Map
}

func newCache() *cacheT {
	c := new(cacheT)
	c.m = new(sync.Map)
	return c
}

func (c *cacheT) get(name string) (messaging.Content, error) {
	v, ok := c.m.Load(name)
	if !ok {
		return messaging.Content{}, errors.New(fmt.Sprintf("resource [%v] not found", name))
	}
	if v1, ok1 := v.(messaging.Content); ok1 {
		return v1, nil
	}
	return messaging.Content{}, nil
}

func (c *cacheT) put(name string, ct messaging.Content) {
	c.m.Store(name, ct)
}
