package referee

import (
    "fmt"
)

type ServerConfig struct {
    RethinkHost string
}

type Server struct {
    config  *ServerConfig
    db      *Db
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

    return nil
}

func (serv *Server) Run() {
    fmt.Println("Starting referee server")
}
