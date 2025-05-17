package repository

import (
	"github.com/behavioral-ai/core/messaging"
	"sync"
)

// mapT - constructor map
type mapT struct {
	m *sync.Map
}

// newMap - create a new map
func newMap() *mapT {
	c := new(mapT)
	c.m = new(sync.Map)
	return c
}

func (c *mapT) get(name string) func() messaging.NewAgent {
	v, ok := c.m.Load(name)
	if !ok {
		return nil
	}
	if v1, ok1 := v.(func() messaging.NewAgent); ok1 {
		return v1
	}
	return nil
}

func (c *mapT) put(name string, fn messaging.NewAgent) {
	if name == "" || fn == nil {
		return
	}
	c.m.Store(name, fn)
}
