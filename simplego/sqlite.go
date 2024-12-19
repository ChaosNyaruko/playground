package main

import (
	"database/sql"
	"fmt"
	_ "github.com/ncruces/go-sqlite3/driver
)

var db *sql.DB

func main() {
	// Get a database handle.
	var err error
	db, err = sql.Open("sqlite", )
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
