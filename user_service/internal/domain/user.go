package domain

import "golang.org/x/crypto/bcrypt"

type (
	UserDomain struct {
		uuid         string
		login        string
		passwordHash []byte
	}

	User interface {
		GetUUID() string
		GetLogin() string
		GetPassword() []byte
		SetUUID(uuid string)
		SetLogin(login string)
		SetPassword(password string) error
	}
)

func NewUser() User {
	return &UserDomain{}
}

func (u *UserDomain) GetUUID() string {
	return u.uuid
}

func (u *UserDomain) GetLogin() string {
	return u.login
}

func (u *UserDomain) GetPassword() []byte {
	return u.passwordHash
}

func (u *UserDomain) SetUUID(uuid string) {
	u.uuid = uuid
}

func (u *UserDomain) SetLogin(login string) {
	u.login = login
}

func (u *UserDomain) SetPassword(password string) error {
	passwordHash, err := u.hashingPassword(password)
	if err != nil {
		return err
	}
	u.passwordHash = passwordHash
	return nil
}

func (u *UserDomain) hashingPassword(password string) ([]byte, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	return passwordHash, nil
}
