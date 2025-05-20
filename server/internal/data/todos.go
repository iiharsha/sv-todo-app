package data

import (
	"time"
)

type Todo struct {
	ID          int64         `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description,omitempty"`
	Completed   bool          `json:"completed"`
	CreatedAt   time.Time     `json:"created_at"`
	DueDate     time.Time     `json:"due_date,omitzero"`
	Priority    PriorityLevel `json:"priority"`
	Tags        []string      `json:"tags,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitzero"`
	Version     int32         `json:"version"`
}
