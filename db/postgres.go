package db

import (
	"context"
	"database/sql"

	"github.com/ryanyogan/cqrs/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertShout(ctx context.Context, shout schema.Shout) error {
	_, err := r.db.Exec("INSERT INTO shouts(id, body, created_at) VALUES($1, $2, $3)",
		shout.ID, shout.Body, shout.CreatedAt)

	return err
}

func (r *PostgresRepository) ListShouts(ctx context.Context, skip uint64, take uint64) ([]schema.Shout, error) {
	rows, err := r.db.Query("SELECT * FROM shouts ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shouts := []schema.Shout{}
	for rows.Next() {
		shout := schema.Shout{}
		if err = rows.Scan(&shout.ID, &shout.Body, &shout.CreatedAt); err == nil {
			shouts = append(shouts, shout)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return shouts, nil
}
