package service

import (
	"auth_service/internal/datasource"
	"auth_service/internal/domain"
	"auth_service/internal/dto"
	"auth_service/pkg/jwt"
	"auth_service/pkg/uuid"

	"github.com/gin-gonic/gin"
)

type authServiceImpl struct {
	Storage datasource.Storage
}

type AuthService interface {
	CreateUser(ctx *gin.Context, userReg dto.DtoRegUser) (string, error)
	AuthUserByLogin(ctx *gin.Context, userAuth dto.DtoAuthUserLogin) (string, error)
}

func NewService(store datasource.Storage) AuthService {
	return &authServiceImpl{
		Storage: store,
	}
}

func (a *authServiceImpl) CreateUser(ctx *gin.Context, userReg dto.DtoRegUser) (string, error) {
	user := domain.NewUser()
	user.SetUUID(uuid.GenerateUUID())
	user.SetLogin(userReg.Login)
	user.SetEmail(userReg.Email)
	user.SetPhone(userReg.Phone)
	user.SetPassword(userReg.Password)

	if err := user.ValidateEmail(); err != nil {
		return "", err
	}
	if err := user.ValidateLogin(); err != nil {
		return "", err
	}
	if err := user.ValidatePhone(); err != nil {
		return "", err
	}
	if err := user.ValidatePassword(); err != nil {
		return "", err
	}
	if err := user.HashingPassword(); err != nil {
		return "", err
	}

	userOut := dto.DtoRegUserDB{
		UUID:         user.GetUUID(),
		Login:        user.GetLogin(),
		Email:        user.GetEmail(),
		Phone:        user.GetPhone(),
		PasswordHash: user.GetPasswordHash(),
	}

	uuid, err := a.Storage.CreateUser(ctx, userOut)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (a *authServiceImpl) AuthUserByLogin(ctx *gin.Context, userAuth dto.DtoAuthUserLogin) (string, error) {
	user := domain.NewUser()
	user.SetLogin(userAuth.Login)
	if err := user.ValidateLogin(); err != nil {
		return "", err
	}
	user.SetPassword(userAuth.Password)
	if err := user.ValidatePassword(); err != nil {
		return "", err
	}
	userIn, err := a.Storage.GetUserByLogin(ctx, userAuth.Login)
	if err != nil {
		return "", err
	}
	user.SetPasswordHash(userIn.PasswordHash)
	if err := user.ApprovePassword(); err != nil {
		return "", err
	}
	userOut := dto.DtoUserToToken{
		UUID:  userIn.UUID,
		Login: userIn.Login,
	}
	token, err := jwt.GenerateToken(userOut)
	if err != nil {
		return "", err
	}
	return token, nil
}
