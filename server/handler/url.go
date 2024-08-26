package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/aditya-amlesh-jha/url-shortener/models"
	"github.com/aditya-amlesh-jha/url-shortener/utils"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type URLHandler struct {
	db          *sql.DB
	redisClient *redis.Client
}

func expiryTime() time.Duration {
	expiration := 30 * 24 * 60 * 60
	expiryDuration := time.Duration(expiration) * time.Second
	return expiryDuration
}

func GetNewURLHandler(db *sql.DB, redisClient *redis.Client) *URLHandler {
	return &URLHandler{
		db:          db,
		redisClient: redisClient,
	}
}

func (h *URLHandler) ShortURL(w http.ResponseWriter, r *http.Request) {
	var req map[string]string

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	longURL := req["long_url"]
	shortURL := utils.GenerateShortURL(longURL)

	_, err := models.InsertURL(h.db, shortURL, longURL)

	log.Printf("Value of shortUrl is %v :: ", shortURL)

	if err != nil {
		http.Error(w, "Failed to store URL", http.StatusInternalServerError)
		return
	}

	h.redisClient.Set(ctx, shortURL, longURL, expiryTime())
}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/redirect/"):]

	longURL, err := h.redisClient.Get(ctx, shortURL).Result()

	if err == redis.Nil {

		url, err := models.GetUrlByShort(h.db, shortURL)

		if err != nil {
			http.Error(w, "URL Not Found", http.StatusNotFound)
			return
		}

		longURL = url.LongURL

		h.redisClient.Set(ctx, shortURL, longURL, expiryTime())

	} else if err != nil {
		http.Error(w, "Failed to retrieve URL", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
