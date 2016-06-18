package bbcsport

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/BarryMcAuley/golden_goal/referee/event"
	"github.com/BarryMcAuley/golden_goal/referee/provider"
	"github.com/PuerkitoBio/goquery"
)

type newMatchProvider struct {
	provider.BaseProvider
}

func (p *newMatchProvider) MainLoop() {
	for {
		events, err := p.scrapeMatchList()
		if err == nil {
			for _, e := range events {
				p.SendEvent(e)
			}
		} else {
			fmt.Println("Error scraping BBC website: " + err.Error())
		}

		time.Sleep(10 * time.Second)
	}
}

func (p *newMatchProvider) scrapeMatchList() ([]*event.Event, error) {
	events := []*event.Event{}

	resp, err := http.Get("http://www.bbc.co.uk/sport/football/live-scores")
	if err != nil {
		return events, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return events, err
	}

	doc.Find(".match-score").Each(func(i int, s *goquery.Selection) {
		home := s.Find(".team-home").Text()
		if len(home) < 1 {
			return
		}

		away := s.Find(".team-away").Text()
		if len(away) < 1 {
			return
		}

		elapsed := s.Find(".elapsed-time").Text()
		if len(elapsed) < 1 || !strings.Contains(strings.ToLower(elapsed), "mins") {
			return
		}

		events = append(events, event.NewMatchEvent(home, away))
	})

	return events, nil
}
