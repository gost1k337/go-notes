package service

import (
	"context"
	"tg_bot/internal/entity"
	"tg_bot/internal/repository"
	"tg_bot/pkg/lib"
	"time"
)

type Note interface {
	GetAll(ctx context.Context) ([]*entity.Note, error)
	Create(ctx context.Context, title, description string, remindsAt []time.Time) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, title, description string, remindsAt []time.Time) (*entity.Note, error)
}

type Services struct {
	Note
}

type ServicesDependencies struct {
	Repos *repository.Repositories
}

func NewServices(deps ServicesDependencies, logger lib.Logger) *Services {
	return &Services{
		Note: NewNoteService(deps.Repos.Note, logger),
	}
}
