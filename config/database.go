package config

import "database/sql"

type DatabaseInterface interface {
	Open() *sql.DB
}

type Database struct {
	Driver string
	DSN    string
}

func NewDatabase(driver, dsn string) *Database {
	return &Database{
		Driver: driver,
		DSN:    dsn,
	}
}

func (db *Database) Open() *sql.DB {
	dbConn, err := sql.Open(db.Driver, db.DSN)
	if err != nil {
		panic(err)
	}
	return dbConn
}
