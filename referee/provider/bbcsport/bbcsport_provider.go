package bbcsport

import (
	"github.com/BarryMcAuley/golden_goal/referee/event"
	"github.com/BarryMcAuley/golden_goal/referee/provider"
)

// Provider BBC Sport website scrape data provider
type Provider struct {
	provider.BaseProvider
	internalChan     chan *event.Event
	newMatchProvider *NewMatchProvider
}

// Initialise Initialises provider and starts provider main loop
func (p *Provider) Initialise() error {
	err := p.BaseProvider.Initialise(nil)
	if err != nil {
		return err
	}

	p.internalChan = make(chan *event.Event)

	go p.MainLoop()

	newMatchProvider := &NewMatchProvider{}
	newMatchProvider.Initialise(p.internalChan)
	go newMatchProvider.MainLoop()

	return nil
}

// GetID Returns the provider ID "BBCSportProvider"
func (p *Provider) GetID() string {
	return "BBCSportProvider"
}

// MainLoop Main provider loop
func (p *Provider) MainLoop() {
	for {
		event := <-p.internalChan
		p.SendEvent(event)
	}
}
