package repository

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

var (
	msg = std.NewSyncMap[string, *messaging.Message]()
)

// LoadMessage - get a message
func LoadMessage(name string) *messaging.Message {
	m, ok := msg.Load(name)
	if !ok {
		return nil
	}
	return m
}

// StoreMessage - store a message
func StoreMessage(m *messaging.Message) {
	msg.Store(m.Name, m)
}

/*
// ModifyMessage - modify a message
func ModifyMessage(m *messaging.Message) {
	msg.modify(m)
}

// DeleteMessage - delete a message
func DeleteMessage(name string) {
	msg.delete(name)
}


*/
