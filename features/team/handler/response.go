package handler

import "time"

type TeamResponse struct {
	ID        uint      `json:"id"`
	TeamName  string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
