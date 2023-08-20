package postgres

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	*sql.DB
	dsn string
}

func New(dsn string) (*Postgres, error) {
	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	d := &Postgres{
		conn,
		dsn,
	}

	return d, nil
}
