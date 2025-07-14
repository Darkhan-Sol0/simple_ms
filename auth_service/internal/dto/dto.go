package dto

type DtoAuthUser struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type DtoAuthUserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type DtoAuthUserEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DtoAuthUserPhone struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type DtoRegUserFromWeb struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type DtoRegUserToDb struct {
	UUID         string `json:"uuid"`
	Login        string `json:"login"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash []byte `json:"passwordhash"`
	Role         int    `json:"id_role"`
}

type DtoUserFromDb struct {
	UUID         string `json:"uuid"`
	Login        string `json:"login"`
	PasswordHash []byte `json:"passwordhash"`
	Role         string `json:"id_role"`
}

type DtoUserToToken struct {
	UUID  string `json:"uuid"`
	Login string `json:"login"`
	Role  string `json:"id_role"`
}

type DtoTokenChecker struct {
	Token string `json:"token"`
}

type DtoUserFromTokenToWeb struct {
	UUID string `json:"uuid"`
	Role string `json:"id_role"`
}
