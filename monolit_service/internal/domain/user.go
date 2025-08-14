package domain

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	regexEmail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regexLogin = `^[a-zA-Z0-9]{3,20}$`
)

type (
	UserDomain struct {
		uuid         string
		login        string
		email        string
		passwordHash []byte
	}

	User interface {
		GetUUID() string
		GetLogin() string
		GetEmail() string
		GetPassword() []byte
		SetUUID(uuid string)
		SetLogin(login string) error
		SetEmail(email string) error
		SetPassword(password string) error
	}
)

func NewUser() User {
	return &UserDomain{}
}

func (u *UserDomain) hashingPassword(password string) ([]byte, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	return passwordHash, nil
}

func (u *UserDomain) validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("error: empty email. Need format (text@mail.domen)")
	}
	re := regexp.MustCompile(regexEmail)

	if !re.MatchString(email) {
		return fmt.Errorf("error: invalid email. Need format (text@mail.domen)")
	}
	return nil
}

func (u *UserDomain) validateLogin(login string) error {
	re := regexp.MustCompile(regexLogin)
	if !re.MatchString(login) {
		return fmt.Errorf("error: invalid login. (only words and numeric, large 3-20 symbols)")
	}
	return nil
}

func (u *UserDomain) SetUUID(uuid string) {
	u.uuid = uuid
}

func (u *UserDomain) SetLogin(login string) error {
	err := u.validateLogin(login)
	if err != nil {
		return err
	}
	u.login = login
	return nil
}

func (u *UserDomain) SetEmail(email string) error {
	err := u.validateEmail(email)
	if err != nil {
		return err
	}
	u.email = email
	return nil
}

func (u *UserDomain) SetPassword(password string) error {
	passwordHash, err := u.hashingPassword(password)
	if err != nil {
		return err
	}
	u.passwordHash = passwordHash
	return nil
}

func (u *UserDomain) GetUUID() string {
	return u.uuid
}

func (u *UserDomain) GetLogin() string {
	return u.login
}

func (u *UserDomain) GetEmail() string {
	return u.email
}

func (u *UserDomain) GetPassword() []byte {
	return u.passwordHash
}
