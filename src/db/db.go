package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var (
	Conn         *sql.DB
	err        error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "connelblaze"
	password = ""
	dbname   = "School_work"
)

//run database connection
func handleDbConnection(val string) {
	// Create an sql.DB and check for errors
	Conn, err = sql.Open("postgres", val)
	if err != nil {
		panic(err.Error())
	}

	// Test the connection to the database
	err = Conn.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected!")
}

//initialize database connection
func DBConnect()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	handleDbConnection(psqlInfo)
}
