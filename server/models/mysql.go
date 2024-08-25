package models

import (
	"database/sql"
	"log"

	"github.com/aditya-amlesh-jha/url-shortener/server/config"
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
