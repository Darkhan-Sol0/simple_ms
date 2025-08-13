package web

import (
	"auth_service/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Responce struct {
	Data any `json:"data"`
}

func (r *routingConfig) PostNewUser(ctx echo.Context) error {
	var user dto.RegUserFromWeb
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Responce{Data: "error"})
	}
	if err := r.service.AddUser(ctx, user); err != nil {
		return ctx.JSON(http.StatusBadRequest, Responce{Data: "error"})
	}
	return ctx.JSON(http.StatusOK, Responce{Data: "OK"})
}

func (r *routingConfig) GetUsersList(ctx echo.Context) error {
	user, err := r.service.GetList(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Responce{Data: "error"})
	}
	return ctx.JSON(http.StatusOK, Responce{Data: user})
}

func (r *routingConfig) GetUserByUuid(ctx echo.Context) error {
	var userUuid dto.GetUserUUIDFromWeb
	err := ctx.Bind(&userUuid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Responce{Data: "error 1"})
	}
	user, err := r.service.GetUserByUuid(ctx, userUuid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Responce{Data: "error 2"})
	}
	return ctx.JSON(http.StatusOK, user)
}
