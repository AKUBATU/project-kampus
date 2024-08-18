package config 

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "samuel:root@/go_tugas")
	if err != nil {
		panic(err)
	}

	log.Println("Database Mulai Terkoneksi nihh, Tungguin yaa")
	DB = db
}