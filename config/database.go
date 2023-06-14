package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:tekaje@/db_bloglrvl1?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("database connected")
	DB = db
}