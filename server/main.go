package main

import (
	"log"
	"net/http"

	"github.com/aditya-amlesh-jha/url-shortener/config"
	"github.com/aditya-amlesh-jha/url-shortener/handler"
	"github.com/aditya-amlesh-jha/url-shortener/models"
)

func main() {

	// Load config
	cfg := config.LoadConfig()

	// Init MySQL and redis
	db := models.InitMySQL(cfg)
	redisClient := models.InitRedis(cfg)

	defer db.Close()
	defer redisClient.Close()

	// Get a url Handler
	urlHandler := handler.GetNewURLHandler(db, redisClient)

	http.HandleFunc("/shorten", urlHandler.ShortURL)
	http.HandleFunc("/redirect/", urlHandler.RedirectURL)

	log.Printf("Starting server on %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(cfg.ServerPort, nil))
}
