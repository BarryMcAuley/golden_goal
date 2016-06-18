package provider

import "github.com/BarryMcAuley/golden_goal/referee/event"

// Provider Interface for all server data-providers
type Provider interface {
	Initialise() error
	GetID() string
	GetEventChannel() chan *event.Event
}

// BaseProvider Reference implementation of the Provider interface with additional
// methods to aid in the construction of data providers. It is expected that all
// concrete data-providers will embed this struct.
type BaseProvider struct {
	eventChan *SafeEventChannel
}

// Initialise Initialises common provider channels and locks
func (p *BaseProvider) Initialise(ch *SafeEventChannel) error {
	if ch == nil {
		ch = NewChannel()
	}

	p.eventChan = ch
	return nil
}

// GetID Returns the ID of the provider
func (p *BaseProvider) GetID() string {
	return "BaseProvider"
}

// GetEventChannel Returns the event channel in use by this provider
func (p *BaseProvider) GetEventChannel() chan *event.Event {
	return p.eventChan.getChannel()
}

// SendEvent Sends an event in a thread-safe manner over the event channel
func (p *BaseProvider) SendEvent(event *event.Event) {
	p.eventChan.SendEvent(event)
}
