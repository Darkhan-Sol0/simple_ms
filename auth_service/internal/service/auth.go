package service

import "gateway/internal/datasource"

type authServiceImpl struct {
	Storage datasource.Storage
}

type AuthService interface {
}

func NewService(store datasource.Storage) AuthService {
	return &authServiceImpl{
		Storage: store,
	}
}
