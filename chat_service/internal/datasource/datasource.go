package datasource

import "chat_service/infrastructure/database"

type Storage interface {
}

type Repository struct {
	Client database.Client
}

func NewRepository(client database.Client) Storage {
	return &Repository{
		Client: client,
	}
}
