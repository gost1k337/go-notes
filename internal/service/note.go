package service

import (
	"context"
	"fmt"
	"tg_bot/internal/entity"
	"tg_bot/internal/repository"
	"tg_bot/internal/repository/repoerrors"
	"tg_bot/pkg/lib"
	"time"
)

type NoteService struct {
	noteRepo repository.Note
	log      lib.Logger
}

func NewNoteService(r repository.Note, logger lib.Logger) *NoteService {
	return &NoteService{
		noteRepo: r,
		log:      logger,
	}
}

func (s NoteService) GetAll(ctx context.Context) ([]*entity.Note, error) {
	return s.noteRepo.GetAll(ctx)
}

func (s NoteService) Create(ctx context.Context, title, description string, remindsAt []time.Time) error {
	return s.noteRepo.Create(ctx, title, description, remindsAt)
}

func (s NoteService) Delete(ctx context.Context, id int) error {
	err := s.noteRepo.Delete(ctx, id)
	if err != nil {
		if err == repoerrors.ErrNotFound {
			return ErrNoteNotFound
		}
		return fmt.Errorf("err: %w", err)
	}
	return nil
}

func (s NoteService) Update(ctx context.Context, id int, title, description string, remindsAt []time.Time) (*entity.Note, error) {
	note, err := s.noteRepo.Update(ctx, id, title, description, remindsAt)
	if err != nil {
		if err == repoerrors.ErrNotFound {
			return nil, ErrNoteNotFound
		}
		return nil, fmt.Errorf("err: %w", err)
	}
	return note, nil
}
