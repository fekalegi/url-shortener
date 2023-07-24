package shortener

import (
	"time"
)

type Status string

// Link represents a shortened URL entry in the system
type Link struct {
	OriginalURL string     `json:"original_url"`
	ShortURL    string     `json:"short_url"`
	Clicks      int        `json:"clicks"`
	ExpireAt    *time.Time `json:"-"`
}

func (l *Link) GetTimeDiffInSecs() int32 {
	now := time.Now()
	sub := l.ExpireAt.Sub(now)
	return int32(sub.Seconds())
}
