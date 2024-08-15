package models

import "time"

type Message struct {
	ID          int       `json:"id"`
	Content     string    `json:"content"`
	Processed   bool      `json:"processed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description *string   `json:"description,omitempty"`
}
