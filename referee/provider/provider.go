package provider

import (
	"sync"

	"github.com/BarryMcAuley/golden_goal/referee/event"
)

type Provider interface {
	Initialise() error
	GetId() string
	GetEventChannel() chan *event.Event
}

type BaseProvider struct {
	eventSendLock *sync.Mutex
	eventChan     chan *event.Event
}

func (p *BaseProvider) Initialise() error {
	p.eventChan = make(chan *event.Event)
	p.eventSendLock = &sync.Mutex{}
	return nil
}

func (p *BaseProvider) GetId() string {
	return "BaseProvider"
}

func (p *BaseProvider) GetEventChannel() chan *event.Event {
	return p.eventChan
}

func (p *BaseProvider) SendEvent(event *event.Event) {
	p.eventSendLock.Lock()

	select {
	case p.eventChan <- event:
	default:
	}

	p.eventSendLock.Unlock()
}
