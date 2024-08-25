package models

import (
	"database/sql"
	"log"

	"github.com/aditya-amlesh-jha/url-shortener/config"
	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL(cfg *config.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.MySQLURI)

	if err != nil {
		log.Fatalf("Failed to connect to mysql :: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping mysql :: %v", err)
	}

	return db
}

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY AUTO_INCREMENT, short_url VARCHAR(255), long_url VARCHAR(255))")

	if err != nil {
		log.Fatalf("Failed to create table users %v :: ", err)
	} else{
		log.Print("Table created successfully!")
	}
}