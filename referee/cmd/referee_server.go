package main

import (
    "flag"
    "fmt"
    "os"

    ref "github.com/BarryMcAuley/golden_goal/referee"
)

func main() {
    var dbHost = flag.String("dbhost", "localhost", "Host for RethinkDB server")
    flag.Parse()

    config := ref.ServerConfig{
        RethinkHost: *dbHost,
    }

    server := ref.NewServer(&config)

    err := server.Initialise()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error initialising server: " + err.Error())
        os.Exit(1)
    }

    server.Run()
}
