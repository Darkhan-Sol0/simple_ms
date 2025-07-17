package dto

import "time"

type DtoRegUserLogin struct {
	Login    string `json:"login"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DtoAuthUser struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type DtoUserAuthToken struct {
	UUID string `json:"uuid"`
	Role string `json:"id_role"`
}

type DtoUserUpdateToUser struct {
	UUID        string            `json:"uuid" db:"uuid"`
	Name        string            `json:"name" db:"name"`
	Description string            `json:"description" db:"description"`
	BornDay     time.Time         `json:"bornDay" db:"bornDay"`
	City        string            `json:"city" db:"city"`
	Links       map[string]string `json:"links" db:"links"`
}
