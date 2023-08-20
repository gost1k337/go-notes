package repository

import (
	"context"
	"tg_bot/internal/entity"
	"tg_bot/internal/repository/psql"
	"tg_bot/pkg/lib"
	"tg_bot/pkg/postgres"
	"time"
)

type Note interface {
	Create(ctx context.Context, title, description string, remindsAt []time.Time) error
	GetAll(ctx context.Context) ([]*entity.Note, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (*entity.Note, error)
	Update(ctx context.Context, id int, title, description string, remindsAt []time.Time) (*entity.Note, error)
}

type Repositories struct {
	Note
}

func NewRepositories(pg *postgres.Postgres, logger lib.Logger) *Repositories {
	return &Repositories{
		Note: psql.NewNoteRepo(pg, logger),
	}
}
