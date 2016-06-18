package bbcsport

import "github.com/BarryMcAuley/golden_goal/referee/provider"

// Provider BBC Sport website scrape data provider
type Provider struct {
	provider.BaseProvider
	internalChan     *provider.SafeEventChannel
	newMatchProvider *newMatchProvider
}

// Initialise Initialises provider and starts provider main loop
func (p *Provider) Initialise() error {
	err := p.BaseProvider.Initialise(nil)
	if err != nil {
		return err
	}

	p.internalChan = provider.NewChannel()
	go p.MainLoop()

	p.startNewMatchScraper()
	return nil
}

func (p *Provider) startNewMatchScraper() {
	newMatchProvider := &newMatchProvider{}
	newMatchProvider.Initialise(p.internalChan)
	go newMatchProvider.MainLoop()
}

// GetID Returns the provider ID "BBCSportProvider"
func (p *Provider) GetID() string {
	return "BBCSportProvider"
}

// MainLoop Main provider loop
func (p *Provider) MainLoop() {
	for {
		event := p.internalChan.ReadEvent()
		p.SendEvent(event)
	}
}
