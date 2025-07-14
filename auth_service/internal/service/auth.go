package service

import (
	"auth_service/internal/datasource"
	"auth_service/internal/domain"
	"auth_service/internal/dto"
	"auth_service/pkg/jwt"
	"auth_service/pkg/uuid"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type authServiceImpl struct {
	Storage datasource.Storage
}

type AuthService interface {
	CreateUser(ctx *gin.Context, userReg dto.DtoRegUserFromWeb) (string, error)
	AuthUser(ctx *gin.Context, userAuth dto.DtoAuthUser) (string, error)
	AuthUserByLogin(ctx *gin.Context, userAuth dto.DtoAuthUserLogin) (string, error)
	TokenChecker(ctx *gin.Context, token string) (dto.DtoUserFromTokenToWeb, error)
}

func NewService(store datasource.Storage) AuthService {
	return &authServiceImpl{
		Storage: store,
	}
}

func (a *authServiceImpl) CreateUser(ctx *gin.Context, userReg dto.DtoRegUserFromWeb) (string, error) {
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

	userOut := dto.DtoRegUserToDb{
		UUID:         user.GetUUID(),
		Login:        user.GetLogin(),
		Email:        user.GetEmail(),
		Phone:        user.GetPhone(),
		PasswordHash: user.GetPasswordHash(),
		Role:         user.GetRole(),
	}

	uuid, err := a.Storage.CreateUser(ctx, userOut)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (a *authServiceImpl) AuthUser(ctx *gin.Context, userAuth dto.DtoAuthUser) (token string, err error) {
	log.Println(userAuth)
	if isEmail(userAuth.Identifier) {
		user := dto.DtoAuthUserEmail{Email: userAuth.Identifier, Password: userAuth.Password}
		if token, err = a.AuthUserByEmail(ctx, user); err != nil {
			return "", err
		}
	} else if isPhone(userAuth.Identifier) {
		user := dto.DtoAuthUserPhone{Phone: userAuth.Identifier, Password: userAuth.Password}
		if token, err = a.AuthUserByPhone(ctx, user); err != nil {
			return "", err
		}
	} else if isLogin(userAuth.Identifier) {
		user := dto.DtoAuthUserLogin{Login: userAuth.Identifier, Password: userAuth.Password}
		if token, err = a.AuthUserByLogin(ctx, user); err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("invalid identifier")
	}
	return token, nil
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
		Role:  userIn.Role,
	}
	token, err := jwt.GenerateToken(userOut)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authServiceImpl) AuthUserByEmail(ctx *gin.Context, userAuth dto.DtoAuthUserEmail) (string, error) {
	user := domain.NewUser()
	user.SetEmail(userAuth.Email)
	if err := user.ValidateEmail(); err != nil {
		return "", err
	}
	user.SetPassword(userAuth.Password)
	if err := user.ValidatePassword(); err != nil {
		return "", err
	}
	userIn, err := a.Storage.GetUserByEmail(ctx, userAuth.Email)
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
		Role:  userIn.Role,
	}
	token, err := jwt.GenerateToken(userOut)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authServiceImpl) AuthUserByPhone(ctx *gin.Context, userAuth dto.DtoAuthUserPhone) (string, error) {
	user := domain.NewUser()
	user.SetPhone(userAuth.Phone)
	if err := user.ValidatePhone(); err != nil {
		return "", err
	}
	user.SetPassword(userAuth.Password)
	if err := user.ValidatePassword(); err != nil {
		return "", err
	}
	userIn, err := a.Storage.GetUserByPhone(ctx, userAuth.Phone)
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
		Role:  userIn.Role,
	}
	token, err := jwt.GenerateToken(userOut)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authServiceImpl) TokenChecker(ctx *gin.Context, token string) (dto.DtoUserFromTokenToWeb, error) {
	claim, err := jwt.ParseToken(token)
	if err != nil {
		return dto.DtoUserFromTokenToWeb{}, err
	}
	user, err := a.Storage.GetUserByLogin(ctx, claim.Login)
	if err != nil {
		return dto.DtoUserFromTokenToWeb{}, err
	}

	if user.UUID != claim.UUID {
		return dto.DtoUserFromTokenToWeb{}, fmt.Errorf("invalid token")
	}

	return dto.DtoUserFromTokenToWeb{
		UUID: user.UUID,
		Role: user.Role,
	}, nil
}
