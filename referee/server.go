package referee

import (
    "fmt"
)

type Server struct {
    db  *Db
}

func NewServer() *Server {
    return &Server{}
}

func (serv *Server) Initialise() error {
    db, err := newDatabase()
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
