package dto

import "time"

type DtoUuidUserFromWeb struct {
	UUID string `json:"uuid"`
}

type DtoUserToDb struct {
	UUID        string            `json:"uuid" db:"uuid"`
	Name        string            `json:"name" db:"name"`
	Description string            `json:"description" db:"description"`
	BornDay     time.Time         `json:"bornDay" db:"bornDay"`
	City        string            `json:"city" db:"city"`
	Links       map[string]string `json:"links" db:"links"`
}

type DtoUserFromDb struct {
	UUID        string            `json:"uuid" db:"uuid"`
	Name        string            `json:"name" db:"name"`
	Description string            `json:"description" db:"description"`
	BornDay     time.Time         `json:"bornDay" db:"bornDay"`
	City        string            `json:"city" db:"city"`
	Links       map[string]string `json:"links" db:"links"`
}
