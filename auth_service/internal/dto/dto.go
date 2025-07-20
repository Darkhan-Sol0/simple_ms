package dto

import "time"

type DtoAuthUser struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type DtoAuthUserLogin struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type DtoAuthUserEmail struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type DtoAuthUserPhone struct {
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}

type DtoRegUserFromWeb struct {
	Login    string `json:"login" db:"login"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}

type DtoRegUserToDb struct {
	Login        string `json:"login" db:"login"`
	Email        string `json:"email" db:"email"`
	Phone        string `json:"phone" db:"phone"`
	PasswordHash []byte `json:"passwordhash" db:"passwordhash"`
	Role         string `json:"user_role" db:"user_role"`
}

type DtoUserFromDb struct {
	UUID         string `json:"uuid" db:"uuid"`
	Login        string `json:"login" db:"login"`
	PasswordHash []byte `json:"passwordhash" db:"passwordhash"`
	Role         string `json:"user_role" db:"user_role"`
}

type DtoUserInfoFromDb struct {
	UUID  string `json:"uuid" db:"uuid"`
	Login string `json:"login" db:"login"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
	// PasswordHash []byte `json:"passwordhash" db:"passwordhash"`
	PasswordHash string    `json:"passwordhash" db:"passwordhash"`
	Role         string    `json:"user_role" db:"user_role"`
	DateCreate   time.Time `json:"created_at" db:"created_at"`
	DateUpdate   time.Time `json:"updated_at" db:"updated_at"`
	Active       bool      `json:"is_active" db:"is_active"`
}

type DtoUserToToken struct {
	UUID  string `json:"uuid" db:"uuid"`
	Login string `json:"login" db:"login"`
	Role  string `json:"user_role" db:"user_role"`
}

type DtoTokenChecker struct {
	Token string `json:"token" db:"token"`
}

type DtoUserFromTokenToWeb struct {
	UUID string `json:"user_uuid" db:"user_uuid"`
	Role string `json:"user_role" db:"user_role"`
}
