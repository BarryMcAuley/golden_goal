package referee

import (
	"fmt"

	rethink "gopkg.in/dancannon/gorethink.v2"
)

// Db Database connection data
type Db struct {
	rethink *rethink.Session
}

// DbName Name of RethinkDB database used globally
const DbName = "referee"

func newDatabase(config *ServerConfig) (*Db, error) {
	db := Db{}

	dbSession, err := rethink.Connect(rethink.ConnectOpts{
		Address: config.RethinkHost + ":28015",
	})
	if err != nil {
		return nil, err
	}

	db.rethink = dbSession
	return &db, nil
}

func (db *Db) initialiseDatabase() error {
	exists, err := db.hasDatabase(DbName)
	if err != nil {
		return err
	} else if !exists {
		db.createDatabase(DbName)
	}

	return nil
}

func (db Db) hasDatabase(dbName string) (bool, error) {
	curs, err := rethink.DBList().Run(db.rethink)
	if err != nil {
		return false, err
	}

	var dbs []string
	err = curs.All(&dbs)
	if err != nil {
		return false, err
	}

	for _, db := range dbs {
		if db == dbName {
			return true, nil
		}
	}

	return false, nil
}

func (db *Db) createDatabase(dbName string) error {
	fmt.Println("No refereee DB found, creating new database")

	if _, err := rethink.DBCreate(dbName).Run(db.rethink); err != nil {
		return err
	}

	return nil
}
