package db

/*
Stores the database functions
*/

import (
	"log"
	"os"

	"github.com/pbillerot/graduel/config"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
)

// openDB ouverture de la base de données
func openDb() *godb.DB {
	values, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := godb.Open(sqlite.Adapter, values.SqlitePath)
	if err != nil {
		log.Fatal(err)
	}
	db.SetLogger(log.New(os.Stderr, "db", 0))
	return db
}
