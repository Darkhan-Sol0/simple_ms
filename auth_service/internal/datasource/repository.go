package datasource

import (
	"auth_service/infrastructure/database"
	"auth_service/internal/dto"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	CreateUser(ctx *gin.Context, user dto.DtoRegUserToDb) (string, error)
	GetUserByLogin(ctx *gin.Context, login string) (dto.DtoUserFromDbToWeb, error)
	GetUserByEmail(ctx *gin.Context, email string) (dto.DtoUserFromDbToWeb, error)
	GetUserByPhone(ctx *gin.Context, phone string) (dto.DtoUserFromDbToWeb, error)
}

type Repository struct {
	Client database.Client
}

func NewRepository(client database.Client) Storage {
	return &Repository{
		Client: client,
	}
}

func (r *Repository) CreateUser(ctx *gin.Context, user dto.DtoRegUserToDb) (string, error) {
	query := `INSERT INTO users (
		uuid, 
		login, 
		email, 
		phone, 
		password
		) VALUES ($1, $2, $3, $4, $5) RETURNING uuid`
	var uuid string
	err := r.Client.QueryRow(ctx, query, user.UUID, user.Login, user.Email, user.Phone, user.PasswordHash).Scan(&uuid)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}
	return uuid, nil
}

func (r *Repository) GetUserByLogin(ctx *gin.Context, login string) (dto.DtoUserFromDbToWeb, error) {
	guery := `SELECT uuid, login, password FROM users WHERE login = $1`
	var user dto.DtoUserFromDbToWeb
	row := r.Client.QueryRow(ctx, guery, login)
	if err := row.Scan(&user.UUID, &user.Login, &user.PasswordHash); err != nil {
		return dto.DtoUserFromDbToWeb{}, fmt.Errorf("failed to scan user: %w", err)
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(ctx *gin.Context, email string) (dto.DtoUserFromDbToWeb, error) {
	guery := `SELECT login, email, password FROM users WHERE email = $1`
	var user dto.DtoUserFromDbToWeb
	row := r.Client.QueryRow(ctx, guery, email)
	if err := row.Scan(&user.UUID, &user.Login, &user.PasswordHash); err != nil {
		return dto.DtoUserFromDbToWeb{}, fmt.Errorf("failed to scan user: %w", err)
	}
	return user, nil
}

func (r *Repository) GetUserByPhone(ctx *gin.Context, phone string) (dto.DtoUserFromDbToWeb, error) {
	guery := `SELECT uuid, login, password FROM users WHERE phone = $1`
	var user dto.DtoUserFromDbToWeb
	row := r.Client.QueryRow(ctx, guery, phone)
	if err := row.Scan(&user.UUID, &user.Login, &user.PasswordHash); err != nil {
		return dto.DtoUserFromDbToWeb{}, fmt.Errorf("failed to scan user: %w", err)
	}
	return user, nil
}
