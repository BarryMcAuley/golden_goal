package main

import (
    "fmt"
    "os"

    ref "github.com/BarryMcAuley/golden_goal/referee"
)

func main() {
    server := ref.NewServer()

    err := server.Initialise()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error initialising server: " + err.Error())
        os.Exit(1)
    }

    server.Run()
}
