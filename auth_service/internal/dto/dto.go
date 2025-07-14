package dto

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
	UUID         string `json:"uuid" db:"uuid"`
	Login        string `json:"login" db:"login"`
	Email        string `json:"email" db:"email"`
	Phone        string `json:"phone" db:"phone"`
	PasswordHash []byte `json:"passwordhash" db:"passwordhash"`
	Role         int    `json:"id_role" db:"id_role"`
}

type DtoUserFromDb struct {
	UUID         string `json:"uuid" db:"uuid"`
	Login        string `json:"login" db:"login"`
	PasswordHash []byte `json:"passwordhash" db:"passwordhash"`
	Role         string `json:"id_role" db:"id_role"`
}

type DtoUserInfoFromDb struct {
	UUID  string `json:"uuid" db:"uuid"`
	Login string `json:"login" db:"login"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
	// PasswordHash []byte `json:"passwordhash" db:"passwordhash"`
	PasswordHash string `json:"passwordhash" db:"passwordhash"`
	Role         string `json:"id_role" db:"id_role"`
}

type DtoUserToToken struct {
	UUID  string `json:"uuid" db:"uuid"`
	Login string `json:"login" db:"login"`
	Role  string `json:"role" db:"role"`
}

type DtoTokenChecker struct {
	Token string `json:"token" db:"token"`
}

type DtoUserFromTokenToWeb struct {
	UUID string `json:"uuid" db:"uuid"`
	Role string `json:"role" db:"role"`
}
