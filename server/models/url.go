package models

import "database/sql"

type URL struct {
	ID       int64  `json:"id"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

func InsertURL(db *sql.DB, shortURL string, longURL string) (int64, error) {
	res, err := db.Exec("INSERT into urls (short_url, long_url) VALUES (?, ?)", shortURL, longURL)

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func GetUrlByShort(db *sql.DB, shortURL string) (URL, err) {
	var url URL
	row := db.QueryRow("Select id, short_url, long_url from urls where short_url = ?", shortURL)
	if err := row.Scan(&url.ID, &url.ShortURL, &url.LongURL); err != nil {
		return URL{}, err
	}
	return url, nil
}
