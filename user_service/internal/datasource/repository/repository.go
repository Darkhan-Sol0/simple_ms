package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	Client interface {
		Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
		Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
		Begin(ctx context.Context) (pgx.Tx, error)
		Close()
	}

	Repository struct {
		Client Client
	}

	Storage interface {
		AddUser() (string, error)
		GetListUser() (map[string]string, error)
	}
)

func NewDatabase(client Client) Storage {
	return &Repository{
		Client: client,
	}
}

func (r *Repository) AddUser() (string, error) {
	return "Hello, World!", nil
}

func (r *Repository) GetListUser() (map[string]string, error) {

	return map[string]string{
		"hui0": "hui",
		"hui1": "hui",
		"hui2": "hui",
		"hui3": "hui",
	}, nil
}
