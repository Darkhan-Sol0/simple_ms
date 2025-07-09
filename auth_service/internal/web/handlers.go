package web

import (
	"auth_service/internal/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Main(ctx *gin.Context) {
	sendMessage(ctx, NewResult("Hello, World!", http.StatusOK, nil))
}

func (h *Handler) Erro(ctx *gin.Context) {
	sendMessage(ctx, NewResult("Test!", http.StatusBadRequest, fmt.Errorf("this bad reques")))
}

func (h *Handler) Succes(ctx *gin.Context) {
	sendMessage(ctx, NewResult("This succes request!", http.StatusOK, nil))
}

func (h *Handler) Registaration(ctx *gin.Context) {
	var regUser dto.DtoRegUser
	if err := ctx.ShouldBindJSON(&regUser); err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, err))
		return
	}
	uuid, err := h.Service.CreateUser(ctx, regUser)
	if err != nil {
		sendMessage(ctx, NewResult("invalid create user", http.StatusBadRequest, err))
		return
	}
	sendMessage(ctx, NewResult(uuid, http.StatusCreated, nil))
}

func (h *Handler) Authorization(ctx *gin.Context) {
	var authUser dto.DtoAuthUserLogin
	if err := ctx.ShouldBindJSON(&authUser); err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, err))
		return
	}
	token, err := h.Service.AuthUserByLogin(ctx, authUser)
	if err != nil {
		sendMessage(ctx, NewResult("invalid auth user", http.StatusBadRequest, err))
		return
	}
	sendMessage(ctx, NewResult(token, http.StatusCreated, nil))
}
