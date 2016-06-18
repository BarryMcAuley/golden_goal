package provider

import (
	"sync"

	"github.com/BarryMcAuley/golden_goal/referee/event"
)

type SafeEventChannel struct {
	lock    *sync.Mutex
	channel chan *event.Event
}

func NewChannel() *SafeEventChannel {
	return &SafeEventChannel{
		lock:    &sync.Mutex{},
		channel: make(chan *event.Event),
	}
}

func (ch *SafeEventChannel) Lock() {
	ch.lock.Lock()
}

func (ch *SafeEventChannel) Unlock() {
	ch.lock.Unlock()
}

func (ch *SafeEventChannel) SendEvent(event *event.Event) {
	ch.Lock()
	ch.channel <- event
	ch.Unlock()
}

func (ch *SafeEventChannel) ReadEvent() *event.Event {
	return <-ch.channel
}

func (ch *SafeEventChannel) getChannel() chan *event.Event {
	return ch.channel
}
