package repository

import (
	"github.com/behavioral-ai/core/messaging"
	"sync"
)

// ctorM - constructor map
type ctorM struct {
	m *sync.Map
}

// newCtorMap - create a new agent map
func newCtorMap() *ctorM {
	c := new(ctorM)
	c.m = new(sync.Map)
	return c
}

func (c *ctorM) get(name string) messaging.NewAgent {
	v, ok := c.m.Load(name)
	if !ok {
		return nil
	}
	if v1, ok1 := v.(messaging.NewAgent); ok1 {
		return v1
	}
	return nil
}

func (c *ctorM) store(name string, fn messaging.NewAgent) {
	if name == "" || fn == nil {
		return
	}
	c.m.Store(name, fn)
}
