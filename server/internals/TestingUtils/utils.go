package testingutils

import (
	"database/sql"
	"log"

	"github.com/codico/boilerplate/db"
	_ "github.com/lib/pq"
)

type TestDB struct {
	Conn    *sql.DB
	Queries *db.Queries
}

func (c *TestDB) Close() {
	c.Conn.Close()
}

func (c *TestDB) Cleanup() {
	if _, err := c.Conn.Exec("TRUNCATE users CASCADE;"); err != nil {
		log.Fatalf("cannot cleanup database %s", err.Error())
	}
	c.Conn.Close()
}

func NewTestConnection() TestDB {
	conn := PGConnection()
	return TestDB{
		Conn:    conn,
		Queries: db.New(conn),
	}
}

func PGConnection() *sql.DB {
	dbName := "boilerplate_testing"
	testDsn := "host=localhost port=5432 user=postgres password=postgres dbname=boilerplate_testing sslmode=disable"
	conn, err := sql.Open("postgres", testDsn)
	if err != nil {
		log.Fatalf("cannot connect to created database: %s", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("cannot ping test database: %s", err)
	}
	log.Printf("test database created: %s", dbName)
	return conn
}
