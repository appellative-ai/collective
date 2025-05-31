package repository

import (
	"github.com/behavioral-ai/core/messaging"
	"sync"
)

// messageM - messaging map
type messageM struct {
	m *sync.Map
}

// newMessageMap - create a new agent map
func newMessageMap() *messageM {
	m := new(messageM)
	m.m = new(sync.Map)
	return m
}

func (m *messageM) get(name string) *messaging.Message {
	if name == "" {
		return nil
	}
	v, ok := m.m.Load(name)
	if !ok {
		return nil
	}
	if v1, ok1 := v.(*messaging.Message); ok1 {
		return v1
	}
	return nil
}

func (m *messageM) store(msg *messaging.Message) {
	if msg == nil || msg.Name == "" {
		return
	}
	if _, ok := m.m.Load(msg.Name); ok {
		return
	}
	m.m.Store(msg.Name, msg)
}

func (m *messageM) modify(msg *messaging.Message) {
	if msg == nil || msg.Name == "" {
		return
	}
	if _, ok := m.m.Load(msg.Name); !ok {
		return
	}
	m.m.Store(msg.Name, msg)
}

func (m *messageM) delete(name string) {
	if name == "" {
		return
	}
	m.m.Delete(name)
}
