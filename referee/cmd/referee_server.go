package main

import (
	"flag"
	"os"

	ref "github.com/BarryMcAuley/golden_goal/referee"
	log "github.com/Sirupsen/logrus"
)

func main() {
	var dbHost = flag.String("dbhost", "localhost", "Host for RethinkDB server")
	var logDebug = flag.Bool("debug", false, "Enables debug logging")
	flag.Parse()

	if *logDebug {
		log.SetLevel(log.DebugLevel)
	}

	config := ref.ServerConfig{
		RethinkHost: *dbHost,
	}

	server := ref.NewServer(&config)

	err := server.Initialise()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("Failed to initialise server")

		os.Exit(1)
	}

	server.Run()
}
