package dto

type DtoRegUserLogin struct {
	Login    string `json:"login"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DtoAuthUserLogin struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
