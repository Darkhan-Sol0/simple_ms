package dto

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
