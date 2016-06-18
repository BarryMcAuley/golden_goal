package referee

import (
	"fmt"
	"reflect"

	"github.com/BarryMcAuley/golden_goal/referee/event"
	"github.com/BarryMcAuley/golden_goal/referee/provider"
	"github.com/BarryMcAuley/golden_goal/referee/provider/bbcsport"
)

type ServerConfig struct {
	RethinkHost string
}

type Server struct {
	config    *ServerConfig
	db        *Db
	providers []provider.Provider
	exiting   bool
}

func NewServer(config *ServerConfig) *Server {
	return &Server{config: config}
}

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

func (serv *Server) InitialiseProviders() error {
	serv.providers = []provider.Provider{
		&bbcsport.BBCSportProvider{},
	}

	for _, p := range serv.providers {
		err := p.Initialise()
		if err != nil {
			return err
		}
	}

	return nil
}

func (serv *Server) Run() {
	fmt.Println("Starting referee server")

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

		fmt.Printf("Read from channel %#v: %s\n", chans[chosen], value.String())
	}
}
