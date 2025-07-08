package dto

type DtoAuthUserLogin struct {
	Login    string
	Password string
}

type DtoAuthUserEmail struct {
	Email    string
	Password string
}

type DtoAuthUserPhone struct {
	Phone    string
	Password string
}

type DtoRegUser struct {
	Login    string
	Email    string
	Phone    string
	Password string
}

type DtoRegUserDB struct {
	UUID         string
	Login        string
	Email        string
	Phone        string
	PasswordHash []byte
}

type DtoUserFromDB struct {
	UUID         string
	Login        string
	PasswordHash []byte
}

type DtoUserToToken struct {
	UUID  string
	Login string
}
