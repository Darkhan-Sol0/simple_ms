package datasource

import (
	"auth_service/infrastructure/database"
	"auth_service/internal/dto"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	CreateUser(ctx *gin.Context, user dto.DtoRegUserToDb) (dto.DtoUserUUIDToWeb, error)
	GetUserByLogin(ctx *gin.Context, login string) (dto.DtoUserFromDb, error)
	GetUserByEmail(ctx *gin.Context, email string) (dto.DtoUserFromDb, error)
	GetUserByPhone(ctx *gin.Context, phone string) (dto.DtoUserFromDb, error)
	GetUserInfoList(ctx *gin.Context) ([]dto.DtoUserInfoFromDb, error)
}

type Repository struct {
	Client database.Client
}

func NewRepository(client database.Client) Storage {
	return &Repository{
		Client: client,
	}
}

func (r *Repository) CreateUser(ctx *gin.Context, user dto.DtoRegUserToDb) (dto.DtoUserUUIDToWeb, error) {
	query := `INSERT INTO users (
		login, email, phone, password_hash,	user_role
		) VALUES ($1, $2, $3, $4, $5) RETURNING uuid`
	var uuid dto.DtoUserUUIDToWeb
	err := r.Client.QueryRow(ctx, query, user.Login, user.Email, user.Phone, user.PasswordHash, user.Role).Scan(&uuid.UUID)
	if err != nil {
		return dto.DtoUserUUIDToWeb{}, fmt.Errorf("failed to create user: %w", err)
	}
	return uuid, nil
}

func (r *Repository) GetUserByLogin(ctx *gin.Context, login string) (dto.DtoUserFromDb, error) {
	guery := `SELECT uuid, login, password_hash, user_role 
						FROM users
						WHERE login = $1`
	var user dto.DtoUserFromDb
	row := r.Client.QueryRow(ctx, guery, login)
	if err := row.Scan(&user.UUID, &user.Login, &user.PasswordHash, &user.Role); err != nil {
		return dto.DtoUserFromDb{}, fmt.Errorf("failed to scan user: %w", err)
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(ctx *gin.Context, email string) (dto.DtoUserFromDb, error) {
	guery := `SELECT uuid, login, password_hash, user_role 
						FROM users
						WHERE email = $1`
	var user dto.DtoUserFromDb
	row := r.Client.QueryRow(ctx, guery, email)
	if err := row.Scan(&user.UUID, &user.Login, &user.PasswordHash, &user.Role); err != nil {
		return dto.DtoUserFromDb{}, fmt.Errorf("failed to scan user: %w", err)
	}
	return user, nil
}

func (r *Repository) GetUserByPhone(ctx *gin.Context, phone string) (dto.DtoUserFromDb, error) {
	guery := `SELECT uuid, login, password_hash, user_role 
						FROM users
						WHERE phone = $1`
	var user dto.DtoUserFromDb
	row := r.Client.QueryRow(ctx, guery, phone)
	if err := row.Scan(&user.UUID, &user.Login, &user.PasswordHash, &user.Role); err != nil {
		return dto.DtoUserFromDb{}, fmt.Errorf("failed to scan user: %w", err)
	}
	return user, nil
}

func (r *Repository) GetUserInfoList(ctx *gin.Context) ([]dto.DtoUserInfoFromDb, error) {
	guery := `SELECT uuid, login, email, phone, password_hash, user_role, created_at, updated_at, is_active
						FROM users`
	rows, err := r.Client.Query(ctx, guery)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()
	var users []dto.DtoUserInfoFromDb
	for rows.Next() {
		var user dto.DtoUserInfoFromDb
		if err := rows.Scan(&user.UUID, &user.Login, &user.Email, &user.Phone, &user.PasswordHash, &user.Role, &user.DateCreate, &user.DateUpdate, &user.Active); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return users, nil
}
