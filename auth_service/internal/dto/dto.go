package dto

type (
	RegUserFromWeb struct {
		Login    string `json:"login" db:"login"`
		Password string `json:"password" db:"password"`
	}
	RegUserToDb struct {
		Login    string `json:"login" db:"login"`
		Password []byte `json:"password" db:"password"`
	}

	GetUserFromDb struct {
		UUID  string `json:"uuid" db:"uuid"`
		Login string `json:"login" db:"login"`
	}

	GetUserUUIDFromWeb struct {
		UUID string `json:"uuid" db:"uuid"`
	}
)
