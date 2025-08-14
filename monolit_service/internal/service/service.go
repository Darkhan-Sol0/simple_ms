package service

import (
	"auth_service/internal/datasource/repository"
	"auth_service/internal/domain"
	"auth_service/internal/dto"
	"fmt"

	"github.com/labstack/echo/v4"
)

type (
	ServiceYmpl struct {
		Storage repository.Storage
	}

	Service interface {
		AddUser(ctx echo.Context, user dto.RegUserFromWeb) error
		GetList(ctx echo.Context) ([]dto.GetUserFromDb, error)
		GetUserByUuid(ctx echo.Context, userUuid dto.GetUserUUIDFromWeb) (dto.GetUserFromDb, error)
	}
)

func NewService(storage repository.Storage) Service {
	return &ServiceYmpl{
		Storage: storage,
	}
}

func (s *ServiceYmpl) AddUser(ctx echo.Context, user dto.RegUserFromWeb) error {
	userDomain := domain.NewUser()
	if err := userDomain.SetLogin(user.Login); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := userDomain.SetEmail(user.Email); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := userDomain.SetPassword(user.Password); err != nil {
		return fmt.Errorf("%w", err)
	}
	userOut := dto.RegUserToDb{
		Login:    user.Login,
		Password: userDomain.GetPassword(),
		Email:    userDomain.GetEmail(),
	}
	if err := s.Storage.CreateUser(ctx, userOut); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (s *ServiceYmpl) GetList(ctx echo.Context) ([]dto.GetUserFromDb, error) {
	users, err := s.Storage.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return users, nil
}

func (s *ServiceYmpl) GetUserByUuid(ctx echo.Context, userUuid dto.GetUserUUIDFromWeb) (dto.GetUserFromDb, error) {
	user, err := s.Storage.GetUserByUuid(ctx, userUuid)
	if err != nil {
		return dto.GetUserFromDb{}, err
	}
	return user, nil
}
