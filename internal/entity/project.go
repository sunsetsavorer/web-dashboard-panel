package entity

import "time"

type Project struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	HexColor  string    `json:"hex_color"`
	CreatedAt time.Time `json:"created_at"`
}
