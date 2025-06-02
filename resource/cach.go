package resource

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/collective/exchange"
	"sync"
)

// resolutionKey -
type resolutionKey struct {
	Name     string `json:"name"`
	Fragment string `json:"fragment"`
}

type content struct {
	body exchange.Content
}

type cacheT struct {
	m *sync.Map
}

func newCache() *cacheT {
	c := new(cacheT)
	c.m = new(sync.Map)
	return c
}

func (c *cacheT) get(name, fragment string) (exchange.Content, error) {
	key := resolutionKey{Name: name, Fragment: fragment}
	value, ok := c.m.Load(key)
	if !ok {
		return exchange.Content{}, errors.New(fmt.Sprintf("resource [%v] not found", name))
	}
	if value1, ok1 := value.(content); ok1 {
		return value1.body, nil
	}
	return exchange.Content{}, nil
}

func (c *cacheT) put(name, fragment string, ct exchange.Content) {
	c.m.Store(resolutionKey{Name: name, Fragment: fragment}, content{body: ct})
}
