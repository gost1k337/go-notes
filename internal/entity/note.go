package entity

import (
	"time"
)

type Note struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	RemindsAt   []time.Time `json:"remindsAt"`
	Deleted     bool        `json:"deleted"`
}
