package model

import "time"

type URL struct {
	OriginalURL string        `json:"original_url"`
	ShortID     string        `json:"short_id"`
	ShortURL    string        `json:"short_url"`
	Expiration  time.Duration `json:"expiration"`
}
