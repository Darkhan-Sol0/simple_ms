package dto

type (
	RegUserFromWeb struct {
		Login    string `json:"login" db:"login"`
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}

	RegUserToDb struct {
		Login    string `json:"login" db:"login"`
		Email    string `json:"email" db:"email"`
		Password []byte `json:"password" db:"password"`
	}

	AuthLoginFromWeb struct {
		Login    string `json:"login" db:"login"`
		Password string `json:"password" db:"password"`
	}

	AuthLoginToDb struct {
		Login    string `json:"login" db:"login"`
		Password []byte `json:"password" db:"password"`
	}

	AuthEmailFromWeb struct {
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}

	AuthEmailToDb struct {
		Email    string `json:"email" db:"email"`
		Password []byte `json:"password" db:"password"`
	}

	GetUserFromDb struct {
		UUID  string `json:"uuid" db:"uuid"`
		Email string `json:"email" db:"email"`
		Login string `json:"login" db:"login"`
	}

	GetUserUUIDFromWeb struct {
		UUID string `json:"uuid" db:"uuid"`
	}
)
