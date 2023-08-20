package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"tg_bot/internal/entity"
	"tg_bot/internal/repository/repoerrors"
	"tg_bot/pkg/dates"
	"tg_bot/pkg/lib"
	"tg_bot/pkg/postgres"
	"time"
)

type NoteRepo struct {
	db  *postgres.Postgres
	log lib.Logger
}

func NewNoteRepo(db *postgres.Postgres, logger lib.Logger) *NoteRepo {
	return &NoteRepo{db, logger}
}

func (nr *NoteRepo) Create(ctx context.Context, title, description string, remindsAt []time.Time) error {
	query := `INSERT INTO
			notes (
			    title, description, reminds_at
			) VALUES (
			    $1, $2, $3
			)
			`
	var reminds []interface{}
	for _, remind := range remindsAt {
		reminds = append(reminds, remind)
	}
	_, err := nr.db.ExecContext(ctx, query, title, description, pq.Array(reminds))
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

func (nr *NoteRepo) GetAll(ctx context.Context) ([]*entity.Note, error) {
	query := `SELECT title, description, reminds_at, deleted FROM notes WHERE deleted=false`

	rows, err := nr.db.QueryContext(ctx, query)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return []*entity.Note{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	var notes []*entity.Note
	for rows.Next() {
		note := new(entity.Note)
		var remindsSlice []byte

		if err := rows.Scan(
			&note.Title,
			&note.Description,
			&remindsSlice,
			&note.Deleted,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		reminds, err := dates.BytesToTimeSlice(remindsSlice)
		if err != nil {
			return nil, err
		}

		note.RemindsAt = reminds

		notes = append(notes, note)
	}

	return notes, nil
}

func (nr *NoteRepo) Delete(ctx context.Context, id int) error {
	query := `UPDATE notes SET deleted=true WHERE id=$1`

	note, err := nr.FindById(ctx, id)
	if err != nil {
		return err
	}

	if note.Deleted {
		return repoerrors.ErrNotFound
	}

	_, err = nr.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

func (nr *NoteRepo) FindById(ctx context.Context, id int) (*entity.Note, error) {
	query := `SELECT title, description, reminds_at, deleted FROM notes WHERE id=$1`

	rows, err := nr.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	defer rows.Close()

	rows.Next()

	note := new(entity.Note)
	var remindsSlice []byte

	if err := rows.Scan(
		&note.Title,
		&note.Description,
		&remindsSlice,
		&note.Deleted,
	); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	reminds, err := dates.BytesToTimeSlice(remindsSlice)
	if err != nil {
		return nil, err
	}

	note.RemindsAt = reminds
	return note, nil
}

func (nr *NoteRepo) Update(ctx context.Context, id int, title, description string, remindsAt []time.Time) (*entity.Note, error) {
	note, err := nr.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if note.Deleted {
		return nil, repoerrors.ErrNotFound
	}

	query := `UPDATE notes SET title = $1, description = $2, reminds_at = $3 WHERE id = $4`
	var reminds []interface{}
	for _, remind := range remindsAt {
		reminds = append(reminds, remind)
	}

	_, err = nr.db.ExecContext(ctx, query, title, description, pq.Array(reminds), id)
	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	note, err = nr.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return note, nil
}
