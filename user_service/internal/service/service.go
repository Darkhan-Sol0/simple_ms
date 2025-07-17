package service

import (
	"user_service/internal/datasource"
	"user_service/internal/domain"
	"user_service/internal/dto"

	"github.com/gin-gonic/gin"
)

type userServiceImpl struct {
	Storage datasource.Storage
}

type UserService interface {
	CreateUser(ctx *gin.Context, user dto.DtoUuidUserFromWeb) (string, error)
	GetUser(ctx *gin.Context, userIn dto.DtoUuidUserFromWeb) (dto.DtoUserFromDb, error)
	UpdateUser(ctx *gin.Context, userIn dto.DtoUserToDb) error
}

func NewService(store datasource.Storage) UserService {
	return &userServiceImpl{
		Storage: store,
	}
}

func (u *userServiceImpl) CreateUser(ctx *gin.Context, userIn dto.DtoUuidUserFromWeb) (string, error) {
	user := domain.NewUser()
	user.SetUUID(userIn.UUID)
	userOut := dto.DtoUserToDb{
		UUID:        user.GetUUID(),
		Name:        user.GetName(),
		Description: user.GetDescription(),
		BornDay:     user.GetBornDay(),
		City:        user.GetCity(),
		Links:       user.GetLinks(),
	}
	res, err := u.Storage.CreateUser(ctx, userOut)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (u *userServiceImpl) GetUser(ctx *gin.Context, userIn dto.DtoUuidUserFromWeb) (dto.DtoUserFromDb, error) {
	res, err := u.Storage.GetUserByUUID(ctx, userIn.UUID)
	if err != nil {
		return dto.DtoUserFromDb{}, err
	}
	return res, nil
}

func (u *userServiceImpl) UpdateUser(ctx *gin.Context, userIn dto.DtoUserToDb) error {
	err := u.Storage.UpdateUser(ctx, userIn)
	if err != nil {
		return err
	}
	return nil
}
