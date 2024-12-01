package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	username := "root"           // Username database (default XAMPP)
	password := ""               // Password database (kosong jika default)
	host := "localhost"          // Host database
	port := "3306"               // Port database
	dbname := "inventory_db" // Nama database

	// DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database connected")
}
