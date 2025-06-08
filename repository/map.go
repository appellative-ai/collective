package repository

import (
	"sync"
)

// mapT - constructor map
type mapT[T, U any] struct {
	m *sync.Map
}

// newMap - create a new agent map
func newMap[T, U any]() *mapT[T, U] {
	c := new(mapT[T, U])
	c.m = new(sync.Map)
	return c
}

func (m *mapT[T, U]) load(t T) (u U) {
	v, ok := m.m.Load(t)
	if !ok {
		return u
	}
	if v1, ok1 := v.(U); ok1 {
		return v1
	}
	return u
}

func (m *mapT[T, U]) store(t T, u U) {
	m.m.Store(t, u)
}
