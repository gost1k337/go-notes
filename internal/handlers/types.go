package handlers

import (
	"tg_bot/internal/entity"
	"time"
)

type CreateNoteInput struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	RemindsAt   []time.Time `json:"remindsAt"`
}

type UpdateNoteInput struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	RemindsAt   []time.Time `json:"remindsAt"`
}

type NotesResponse struct {
	Notes []*entity.Note
}

type UpdatedNoteResponse struct {
	Note *entity.Note
}
