package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// db connection test
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)
var testQueries *Queries //*Queries is from db.go, Queries has DBTX interface type "db"
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err := sql.Open(dbDriver, dbSource) //db 연결

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB) //New is from db.go DBTX 타입, *Queries를 반환

	os.Exit(m.Run())

} 