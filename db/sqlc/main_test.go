package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
