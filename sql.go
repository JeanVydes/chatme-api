package main

import (
	"os"
	"fmt"
	"database/sql"

	_ "github.com/lib/pq" // Postgres driver
)

type Database struct {
	conn *sql.DB
	conn_uri string
}

func (db *Database) Connect(core *Core) {
	var err error
	connection_type := os.Getenv("DB_CONNECTION")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_NAME")

	db.conn_uri = fmt.Sprintf("%s://%s:%s@%s:%s/%s", connection_type, user, os.Getenv("DB_PASSWORD"), host, port, dbname)
	db.conn, err = sql.Open("postgres", db.conn_uri)
	if err != nil {
		panic(err)
	}

	err = db.conn.Ping()
	if err != nil {
		panic(err)
	}

	core.db = db
	core.logger.Println("POSTGRES CONNECTED, LOGGED IN AS: " + user)
}

func (db *Database) Disconnect() {
	db.conn.Close()
}

func (db *Database) GetConnection() *sql.DB {
	return db.conn
}

func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

func (db *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.conn.QueryRow(query, args...)
}

func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}