package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/ncruces/go-sqlite3/driver"
	// _ "github.com/ncruces/go-sqlite3/embed"
)

const (
	recordCount = 10000 // Number of records to insert and query
)

func benchmarkSQLite() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS peers (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Insert records
	// start := time.Now()
	// for i := 0; i < recordCount; i++ {
	// 	_, err = db.Exec("INSERT INTO peers (name) VALUES (?)", fmt.Sprintf("Peer %d", i))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// fmt.Printf("SQLite Insert Time: %v\n", time.Since(start))

	// Query records
	start := time.Now()
	rows, err := db.Query("SELECT name FROM peers")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SQLite Query Time 1: %v\n", time.Since(start))
	defer rows.Close()

	var name string
	for rows.Next() {
		rows.Scan(&name)
	}
	fmt.Printf("SQLite Query Time: %v\n", time.Since(start))
}

func benchmarkMySQL() {
	dsn := "wft:1@tcp(10.37.15.26:3306)/mysql" // Update with your MySQL credentials
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS peers (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	if err != nil {
		log.Fatal(err)
	}

	// Insert records
	// start := time.Now()
	// for i := 0; i < recordCount; i++ {
	// 	_, err = db.Exec("INSERT INTO peers (name) VALUES (?)", fmt.Sprintf("Peer %d", i))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// fmt.Printf("MySQL Insert Time: %v\n", time.Since(start))

	// Query records
	start := time.Now()
	rows, err := db.Query("SELECT name FROM peers")
	fmt.Printf("MySQL Query Time 1: %v\n", time.Since(start))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		rows.Scan(&name)
	}
	fmt.Printf("MySQL Query Time: %v\n", time.Since(start))
}

func main() {
	benchmarkSQLite()
	benchmarkMySQL()
}
