package service

import "chat_service/internal/datasource"

type chatServiceImpl struct {
	Storage datasource.Storage
}

type ChatService interface {
}

func NewService(store datasource.Storage) ChatService {
	return &chatServiceImpl{
		Storage: store,
	}
}
