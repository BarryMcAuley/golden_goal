package bbcsport

import (
	"github.com/BarryMcAuley/golden_goal/referee/event"
	"github.com/BarryMcAuley/golden_goal/referee/provider"
)

type BBCSportProvider struct {
	provider.BaseProvider
}

func (p *BBCSportProvider) Initialise() error {
	err := p.BaseProvider.Initialise()
	if err != nil {
		return err
	}

	go p.MainLoop()

	return nil
}

func (p *BBCSportProvider) GetId() string {
	return "BBCSportProvider"
}

func (p *BBCSportProvider) MainLoop() {

	p.SendEvent(&event.Event{})

}
