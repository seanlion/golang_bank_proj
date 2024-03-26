package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) { // testing을 실행하려면 하나의 패키지에서 실행한다?
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn) // 이렇게 new만 해도, db(DbTX)라는걸 알 수 있나? 결국 New(시그니쳐)로 파악해야함.
	os.Exit(m.Run())
}
