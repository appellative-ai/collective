package content

import (
	"errors"
	"fmt"
	"sync"
)

// resolutionKey -
type resolutionKey struct {
	Name     string `json:"name"`
	Fragment string `json:"fragment"`
}

type content struct {
	body Content
}

type cacheT struct {
	m *sync.Map
}

func newCache() *cacheT {
	c := new(cacheT)
	c.m = new(sync.Map)
	return c
}

func (c *cacheT) get(name, fragment string) (Content, error) {
	key := resolutionKey{Name: name, Fragment: fragment}
	value, ok := c.m.Load(key)
	if !ok {
		return Content{}, errors.New(fmt.Sprintf("content [%v] not found", name))
	}
	if value1, ok1 := value.(content); ok1 {
		return value1.body, nil
	}
	return Content{}, nil
}

func (c *cacheT) put(name, fragment string, ct Content) {
	c.m.Store(resolutionKey{Name: name, Fragment: fragment}, content{body: ct})
}
