package resolution

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/core/std"
	"sync"
)

// resolutionKey -
type resolutionKey struct {
	Name     string `json:"name"`
	Fragment string `json:"fragment"`
}

type cacheT struct {
	m *sync.Map
}

func newCache() *cacheT {
	c := new(cacheT)
	c.m = new(sync.Map)
	return c
}

func (c *cacheT) get(name string) (std.Content, error) {
	v, ok := c.m.Load(name)
	if !ok {
		return std.Content{}, errors.New(fmt.Sprintf("resolution [%v] not found", name))
	}
	if v1, ok1 := v.(std.Content); ok1 {
		return v1, nil
	}
	return std.Content{}, nil
}

func (c *cacheT) put(name string, ct std.Content) {
	c.m.Store(name, ct)
}
