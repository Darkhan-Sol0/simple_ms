package web

import (
	"auth_service/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Registaration(ctx *gin.Context) {
	var regUser dto.DtoRegUserFromWeb
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
	var authUser dto.DtoAuthUser
	if err := ctx.ShouldBindJSON(&authUser); err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, err))
		return
	}
	token, err := h.Service.AuthUser(ctx, authUser)
	if err != nil {
		sendMessage(ctx, NewResult("invalid auth user", http.StatusBadRequest, err))
		return
	}
	sendMessage(ctx, NewResult(token, http.StatusCreated, nil))
}

func (h *Handler) CheckAuthorization(ctx *gin.Context) {
	var token string
	if err := ctx.ShouldBindJSON(&token); err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, err))
		return
	}
	userOut, err := h.Service.TokenChecker(ctx, token)
	if err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusUnauthorized, err))
		return
	}
	sendMessage(ctx, NewResult(userOut, http.StatusOK, nil))
}

func (h *Handler) GetUsersList(ctx *gin.Context) {
	userList, err := h.Service.GetUsersList(ctx)
	if err != nil {
		sendMessage(ctx, NewResult("invalid internal service", http.StatusBadRequest, err))
		return
	}
	sendMessage(ctx, NewResult(userList, http.StatusOK, nil))
}
