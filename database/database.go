package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

const (
	server   = "Andrei-PC"
	port     = 1433
	database = "DevNest"
)

var DB *sql.DB

// InitializeDB opens and initializes the database connection.
func InitializeDB() {
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;trusted_connection=yes;", server, port, database)
	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB")
}
