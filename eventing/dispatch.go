package eventing

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

type DispatchItem struct {
	Agent   messaging.Agent
	Channel any
	Event   string
}

type Dispatcher interface {
	Dispatch(agent messaging.Agent, channel any, event string)
}

func Dispatch(agent messaging.Agent, dispatcher Dispatcher, channel any, event string) {
	if dispatcher == nil || agent == nil || channel == nil {
		return
	}
	if ch, ok := channel.(*messaging.Channel); ok {
		dispatcher.Dispatch(agent, "ch:"+ch.Name(), event)
		return
	}
	if t, ok := channel.(*messaging.Ticker); ok {
		dispatcher.Dispatch(agent, "tk:"+t.Name(), event)
	}
}

type traceDispatch struct {
	allEvents bool
	channel   string
	m         map[string]string
}

func (t *traceDispatch) validEvent(event string) bool {
	if t.allEvents {
		return true
	}
	if _, ok := t.m[event]; ok {
		return true
	}
	return false
}

func (t *traceDispatch) validChannel(channel string) bool {
	if t.channel == "" {
		return true
	}
	return t.channel == channel
}

func (t *traceDispatch) Dispatch(agent messaging.Agent, channel any, event string) {
	if !t.validEvent(event) {
		return
	}
	name := "<nil>"
	if ch, ok := channel.(*messaging.Channel); ok {
		name = "ch:" + ch.Name()
		if !t.validChannel(ch.Name()) {
			return
		}
	} else {
		if tk, ok1 := channel.(*messaging.Ticker); ok1 {
			name = "tk:" + tk.Name()
		}
	}
	id := "<nil>"
	if agent != nil {
		id = agent.Uri()
	}
	fmt.Printf("trace -> %v [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), id, name, event)
	//} else {
	//	fmt.Printf("trace -> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), channel, eventing, id, activity)
	//}
}

func NewTraceDispatcher() Dispatcher {
	return NewFilteredTraceDispatcher(nil, "")
}

func NewFilteredTraceDispatcher(events []string, channel string) Dispatcher {
	t := new(traceDispatch)
	if len(events) == 0 {
		t.allEvents = true
	} else {
		t.m = make(map[string]string)
		for _, event := range events {
			t.m[event] = ""
		}
	}
	t.channel = channel
	return t
}

func NewDispatchMessage(agent messaging.Agent, channel any, event string) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, DispatchEvent)
	m.SetContent(ContentTypeDispatch, DispatchItem{Agent: agent, Channel: channel, Event: event})
	return m
}

func DispatchContent(msg *messaging.Message) DispatchItem {
	if msg == nil || msg.ContentType() != ContentTypeDispatch || msg.Body == nil {
		return DispatchItem{}
	}
	if e, ok := msg.Body.(DispatchItem); ok {
		return e
	}
	return DispatchItem{}
}
