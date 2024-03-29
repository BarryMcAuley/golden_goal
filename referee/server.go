package referee

import (
	"reflect"

	log "github.com/Sirupsen/logrus"

	"github.com/BarryMcAuley/golden_goal/referee/event"
	"github.com/BarryMcAuley/golden_goal/referee/provider"
	"github.com/BarryMcAuley/golden_goal/referee/provider/bbcsport"
)

// ServerConfig Server startup configuration data
type ServerConfig struct {
	RethinkHost string
}

// Server Top-level server data-structure
type Server struct {
	config    *ServerConfig
	db        *Db
	providers []provider.Provider
	exiting   bool
}

// NewServer Creates a new referee server from the given config
func NewServer(config *ServerConfig) *Server {
	return &Server{config: config}
}

// Initialise Initialises the server's database connection and data providers
func (serv *Server) Initialise() error {
	db, err := newDatabase(serv.config)
	if err != nil {
		return err
	}

	serv.db = db
	err = db.initialiseDatabase()
	if err != nil {
		return err
	}

	err = serv.InitialiseProviders()
	if err != nil {
		return err
	}

	return nil
}

// InitialiseProviders Initialises the server's  data providers
func (serv *Server) InitialiseProviders() error {
	serv.providers = []provider.Provider{
		&bbcsport.Provider{},
	}

	for _, p := range serv.providers {
		err := p.Initialise()
		if err != nil {
			return err
		}
	}

	return nil
}

// Run Starts execution of the server's main loop
func (serv *Server) Run() {
	log.Info("Starting referee server")

	chans := [](chan *event.Event){}
	for _, prov := range serv.providers {
		chans = append(chans, prov.GetEventChannel())
	}

	cases := make([]reflect.SelectCase, len(chans))
	for i, ch := range chans {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}
	}

	for !serv.exiting {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			cases[chosen].Chan = reflect.ValueOf(nil)
			continue
		}

		matchEvent := *value.Interface().(*event.Event)

		switch matchEvent.EventType {
		case event.EventNewMatch:
			serv.handleNewMatchEvent(&matchEvent)
		}
	}
}

func (serv *Server) handleNewMatchEvent(event *event.Event) {
	if serv.db.hasLiveMatch(event.EventTeamHome, event.EventTeamAway) {
		log.WithFields(log.Fields{
			"event": *event,
		}).Debug("Duplicate new match")
	} else {
		match := newMatch(event.EventTeamHome, event.EventTeamAway)

		log.WithFields(log.Fields{
			"match": match,
		}).Info("New match")

		if err := serv.db.addLiveMatch(match); err != nil {
			log.WithFields(log.Fields{
				"match": match,
				"error": err.Error(),
			}).Error("Failed to add new match to database")
		} else {
			log.WithFields(log.Fields{
				"match": match,
			}).Debug("Failed to add new match to database")

			serv.broadcastToProviders(event, []string{
				event.EventSource,
			})
		}
	}
}

func (serv *Server) broadcastToProviders(event *event.Event, filter []string) {
	for _, p := range serv.providers {
		canSkip := false
		for _, filtered := range filter {
			if p.GetID() == filtered {
				canSkip = true
				break
			}
		}

		if canSkip {
			continue
		}

		ch := p.GetIncomingChannel()
		if ch != nil {
			ch.SendEvent(event)
		}
	}

}
