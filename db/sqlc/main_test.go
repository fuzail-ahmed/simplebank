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

func TestMain(m *testing.M) {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
