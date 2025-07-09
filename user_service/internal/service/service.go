package service

import "user_service/internal/datasource"

type userServiceImpl struct {
	Storage datasource.Storage
}

type UserService interface {
}

func NewService(store datasource.Storage) UserService {
	return &userServiceImpl{
		Storage: store,
	}
}
