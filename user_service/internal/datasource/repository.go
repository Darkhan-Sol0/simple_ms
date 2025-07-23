package datasource

import (
	"encoding/json"
	"fmt"
	"strings"
	"user_service/infrastructure/database"
	"user_service/internal/dto"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	CreateUser(ctx *gin.Context, user dto.DtoUserToDb) (string, error)
	GetUserByUUID(ctx *gin.Context, uuid string) (dto.DtoUserFromDb, error)
	UpdateUser(ctx *gin.Context, user dto.DtoUserToDb) error
	GetUserList(ctx *gin.Context) ([]dto.DtoUserFromDb, error)
}

type Repository struct {
	Client database.Client
}

func NewRepository(client database.Client) Storage {
	return &Repository{
		Client: client,
	}
}

func (r *Repository) CreateUser(ctx *gin.Context, user dto.DtoUserToDb) (string, error) {
	query := `INSERT INTO users (
		uuid, name, description, born_day, city, links
		) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid`
	var uuid string
	linksJSON, err := json.Marshal(user.Links)
	if err != nil {
		return "", fmt.Errorf("failed to marshal links: %w", err)
	}
	err = r.Client.QueryRow(ctx, query, user.UUID, user.Name, user.Description, user.BornDay, user.City, linksJSON).Scan(&uuid)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}
	return uuid, nil
}

func (r *Repository) GetUserByUUID(ctx *gin.Context, uuid string) (dto.DtoUserFromDb, error) {
	guery := `SELECT uuid, name, description, born_day, city, links
						FROM users
						WHERE uuid = $1`
	var user dto.DtoUserFromDb
	var linksJSON []byte
	row := r.Client.QueryRow(ctx, guery, uuid)
	if err := row.Scan(&user.UUID, &user.Name, &user.Description, &user.BornDay, &user.City, &linksJSON); err != nil {
		return dto.DtoUserFromDb{}, fmt.Errorf("failed to scan user: %w", err)
	}
	if err := json.Unmarshal(linksJSON, &user.Links); err != nil {
		return dto.DtoUserFromDb{}, fmt.Errorf("failed to scan user: %w", err)
	}
	return user, nil
}

func (r *Repository) UpdateUser(ctx *gin.Context, user dto.DtoUserToDb) error {
	var setClauses []string
	var args []interface{}
	argID := 1
	if user.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argID))
		args = append(args, user.Name)
		argID++
	}
	if user.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", argID))
		args = append(args, user.Description)
		argID++
	}
	if user.BornDay.IsZero() {
		setClauses = append(setClauses, fmt.Sprintf("born_day = $%d", argID))
		args = append(args, user.BornDay)
		argID++
	}
	if user.City != "" {
		setClauses = append(setClauses, fmt.Sprintf("city = $%d", argID))
		args = append(args, user.City)
		argID++
	}
	if user.Links != nil {
		setClauses = append(setClauses, fmt.Sprintf("links = $%d", argID))
		linksJSON, err := json.Marshal(user.Links)
		if err != nil {
			return fmt.Errorf("failed to marshal links: %w", err)
		}
		args = append(args, linksJSON)
		argID++
	}
	if len(setClauses) == 0 {
		return fmt.Errorf("no fields to update for user with uuid %s", user.UUID)
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE uuid = $%d", strings.Join(setClauses, ", "), argID)
	args = append(args, user.UUID)

	result, err := r.Client.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	if rowsAffected := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("user with uuid %s not found", user.UUID)
	}
	return nil
}

func (r *Repository) GetUserList(ctx *gin.Context) ([]dto.DtoUserFromDb, error) {
	guery := `SELECT uuid, name, description, born_day, city, links
						FROM users`
	rows, err := r.Client.Query(ctx, guery)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()
	var users []dto.DtoUserFromDb
	for rows.Next() {
		var user dto.DtoUserFromDb
		var linksJSON []byte
		if err := rows.Scan(&user.UUID, &user.Name, &user.Description, &user.BornDay, &user.City, &linksJSON); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		if err := json.Unmarshal(linksJSON, &user.Links); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return users, nil
}
