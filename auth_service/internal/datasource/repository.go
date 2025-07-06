package datasource

import "gateway/infrastructure/database"

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
