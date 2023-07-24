package model

import "time"

type LinkRequest struct {
	URL       string     `json:"url" validate:"required"`
	ExpiredAt *time.Time `json:"expired_at"`
}

type SortRequest struct {
	SortBy string `form:"sort_by" validate:"oneof= 'asc' 'desc' ''"`
}
