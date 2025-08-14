package repository

import (
	"auth_service/internal/dto"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
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
		CreateUser(ctx echo.Context, user dto.RegUserToDb) error
		GetUsers(ctx echo.Context) ([]dto.GetUserFromDb, error)
		GetUserByUuid(ctx echo.Context, userUuid dto.GetUserUUIDFromWeb) (dto.GetUserFromDb, error)
	}
)

func NewDatabase(client Client) Storage {
	return &Repository{
		Client: client,
	}
}

func (r *Repository) CreateUser(ctx echo.Context, user dto.RegUserToDb) error {
	query := `
	INSERT INTO users (login, email, password) 
	VALUES ($1, $2, $3)`

	_, err := r.Client.Exec(
		ctx.Request().Context(),
		query,
		user.Login,
		user.Email,
		user.Password,
	)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (r *Repository) GetUsers(ctx echo.Context) ([]dto.GetUserFromDb, error) {
	query := `
	SELECT uuid, login, email
	FROM users`

	rows, err := r.Client.Query(
		ctx.Request().Context(),
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()
	var users []dto.GetUserFromDb

	for rows.Next() {
		var user dto.GetUserFromDb
		if err := rows.Scan(&user.UUID, &user.Login, &user.Email); err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return users, nil
}

func (r *Repository) GetUserByUuid(ctx echo.Context, userUuid dto.GetUserUUIDFromWeb) (dto.GetUserFromDb, error) {
	query := `
	SELECT uuid, login, email
	FROM users
	WHERE uuid = $1`

	row := r.Client.QueryRow(ctx.Request().Context(), query, userUuid.UUID)
	var user dto.GetUserFromDb
	if err := row.Scan(&user.UUID, &user.Login, &user.Email); err != nil {
		return dto.GetUserFromDb{}, fmt.Errorf("error db")
	}
	return user, nil
}
