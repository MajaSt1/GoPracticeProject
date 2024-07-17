package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10) //if we have more request coming to database then it will open max 10 connections
	DB.SetMaxIdleConns(5)  // set value max open connection if nobody uses the database
}
