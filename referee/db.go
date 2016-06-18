package referee

import (
	"fmt"

	rethink "gopkg.in/dancannon/gorethink.v2"
)

// Db Database connection data
type Db struct {
	rethink *rethink.Session
}

const (
	// DbName Name of RethinkDB database used globally
	DbName = "referee"
	// TableLiveMatches Table for live match data
	TableLiveMatches = "LiveMatches"
)

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
		db.rethink.Use(DbName)
		db.createDatabase(DbName)
	}

	db.rethink.Use(DbName)
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

	if err := db.createTables(); err != nil {
		return err
	}

	return nil
}

func (db *Db) createTables() error {
	if _, err := rethink.TableCreate(TableLiveMatches).RunWrite(db.rethink); err != nil {
		return err
	}

	return nil
}

func (db *Db) hasLiveMatch(home string, away string) bool {
	res, err := rethink.Table(TableLiveMatches).Filter(map[string]interface{}{
		"HomeTeam": home,
		"AwayTeam": away,
	}).Run(db.rethink)

	if err != nil {
		fmt.Println("Failed to query db: " + err.Error())
		return false
	}
	defer res.Close()

	var matches []interface{}
	err = res.All(&matches)
	if err != nil {
		fmt.Println("Failed to parse db response: " + err.Error())
		return false
	}

	return len(matches) > 0
}

func (db *Db) addLiveMatch(match *Match) error {
	if _, err := rethink.Table(TableLiveMatches).Insert(*match).RunWrite(db.rethink); err != nil {
		return err
	}

	return nil
}
