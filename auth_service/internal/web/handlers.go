package web

import (
	"auth_service/internal/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Main(ctx *gin.Context) {
	sendMessage(ctx, Result{
		status: http.StatusOK,
		err:    nil,
		data:   "Hello, World!",
	})
}

func (h *Handler) Erro(ctx *gin.Context) {
	sendMessage(ctx, Result{
		status: http.StatusBadRequest,
		err:    fmt.Errorf("this bad reques"),
		data:   "Test!",
	})
}

func (h *Handler) Succes(ctx *gin.Context) {
	sendMessage(ctx, Result{
		status: http.StatusOK,
		err:    nil,
		data:   "This succes request!",
	})
}

func (h *Handler) Registaration(ctx *gin.Context) {
	var regUser dto.DtoRegUser
	if err := ctx.ShouldBindJSON(&regUser); err != nil {
		sendMessage(ctx, Result{
			status: http.StatusBadRequest,
			data:   "invalid request body",
			err:    err})
		return
	}
	uuid, err := h.Service.CreateUser(ctx, regUser)
	if err != nil {
		sendMessage(ctx, Result{
			status: http.StatusBadRequest,
			data:   "invalid create user",
			err:    err})
		return
	}
	sendMessage(ctx, Result{
		status: http.StatusCreated,
		data:   uuid,
		err:    nil})
}

func (h *Handler) Authorization(ctx *gin.Context) {
	var authUser dto.DtoAuthUserLogin
	if err := ctx.ShouldBindJSON(&authUser); err != nil {
		sendMessage(ctx, Result{
			status: http.StatusBadRequest,
			data:   "invalid request body",
			err:    err})
		return
	}
	token, err := h.Service.AuthUserByLogin(ctx, authUser)
	if err != nil {
		sendMessage(ctx, Result{
			status: http.StatusBadRequest,
			data:   "invalid auth user",
			err:    err})
		return
	}
	sendMessage(ctx, Result{
		status: http.StatusCreated,
		data:   token,
		err:    nil})
}
