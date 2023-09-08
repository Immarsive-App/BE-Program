package handler

import "time"

type StatusResponse struct {
	ID         uint      `json:"id"`
	StatusName string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at"`
}
