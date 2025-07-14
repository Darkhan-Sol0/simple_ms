package domain

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	loginRegex = `^[a-zA-Z0-9]{3,20}$`
	phoneRegex = `^\+[0-9]{11}$`

	ADMIN = 1
	USER  = 2
)

type authImpl struct {
	UUID         string
	Login        string
	Email        string
	Phone        string
	Role         int
	Password     string
	PasswordHash []byte
}

type Auth interface {
	ValidateEmail() error
	ValidateLogin() error
	ValidatePhone() error
	ValidatePassword() error
	HashingPassword() error
	ApprovePassword() error

	SetUUID(uuid string)
	SetLogin(login string)
	SetEmail(email string)
	SetPhone(phone string)
	SetPassword(password string)
	SetPasswordHash(passwordHash []byte)
	SetRole(role int)

	GetUUID() string
	GetLogin() string
	GetEmail() string
	GetPhone() string
	GetPassword() string
	GetPasswordHash() []byte
	GetRole() int
}

func NewUser() Auth {
	return &authImpl{
		Role: USER,
	}
}

// Validator

func (a *authImpl) ValidateEmail() error {
	if a.Email == "" {
		// return fmt.Errorf("FAIL: empty email")
		return nil
	}
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(a.Email) {
		return fmt.Errorf("FAIL: invalid email format (text@mail.domen)")
	}
	return nil
}

func (a *authImpl) ValidatePhone() error {
	if a.Phone == "" {
		// return fmt.Errorf("FAIL: empty phone")
		return nil
	}
	re := regexp.MustCompile(phoneRegex)
	if !re.MatchString(a.Phone) {
		return fmt.Errorf("FAIL: invalid phone format (+xxxxxxxxxxx)")
	}
	return nil
}

func (a *authImpl) ValidateLogin() error {
	if a.Login == "" {
		return fmt.Errorf("FAIL: empty login")
	}
	re := regexp.MustCompile(loginRegex)
	if !re.MatchString(a.Login) {
		return fmt.Errorf("FAIL: invalid login format (only words and numeric, large 3-20 symbols)")
	}
	return nil
}

func (a *authImpl) ValidatePassword() error {
	if a.Password == "" {
		return fmt.Errorf("FAIL: empty password")
	}
	return nil
}

func (a *authImpl) HashingPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(a.Password), 10)
	if err != nil {
		return fmt.Errorf("FAIL: error by crypted password")
	}
	a.PasswordHash = passwordHash
	a.Password = ""
	return nil
}

func (a *authImpl) ApprovePassword() error {
	err := bcrypt.CompareHashAndPassword(a.PasswordHash, []byte(a.Password))
	if err != nil {
		return fmt.Errorf("FAIL: invalid password")
	}
	return nil
}

// Setters/Getters

func (a *authImpl) SetUUID(uuid string) {
	a.UUID = uuid
}

func (a *authImpl) SetLogin(login string) {
	a.Login = login
}

func (a *authImpl) SetEmail(email string) {
	a.Email = email
}

func (a *authImpl) SetPhone(phone string) {
	a.Phone = phone
}

func (a *authImpl) SetPassword(password string) {
	a.Password = password
}

func (a *authImpl) SetPasswordHash(passwordHash []byte) {
	a.PasswordHash = passwordHash
}

func (a *authImpl) SetRole(role int) {
	a.Role = role
}

func (a *authImpl) GetUUID() string {
	return a.UUID
}

func (a *authImpl) GetLogin() string {
	return a.Login
}

func (a *authImpl) GetEmail() string {
	return a.Email
}

func (a *authImpl) GetPhone() string {
	return a.Phone
}

func (a *authImpl) GetPassword() string {
	return a.Password
}

func (a *authImpl) GetPasswordHash() []byte {
	return a.PasswordHash
}

func (a *authImpl) GetRole() int {
	return a.Role
}
