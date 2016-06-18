package bbcsport

import (
	"time"

	"github.com/BarryMcAuley/golden_goal/referee/event"
	"github.com/BarryMcAuley/golden_goal/referee/provider"
)

type NewMatchProvider struct {
	provider.BaseProvider
}

func (p *NewMatchProvider) MainLoop() {
	for {
		newMatchEvent := event.NewMatchEvent("Man United", "Liverpool")
		p.SendEvent(newMatchEvent)

		time.Sleep(10 * time.Second)
	}
}
