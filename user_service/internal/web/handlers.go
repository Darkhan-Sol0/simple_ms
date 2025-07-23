package web

import (
	"fmt"
	"net/http"
	"user_service/internal/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	var user dto.DtoUuidUserFromWeb
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, err))
		return
	}
	res, err := h.Service.CreateUser(ctx, user)
	if err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusInternalServerError, err))
		return
	}
	sendMessage(ctx, NewResult(res, http.StatusOK, nil))
}

func (h *Handler) GetUser(ctx *gin.Context) {
	var user dto.DtoUuidUserFromWeb
	user.UUID = ctx.Param("uuid")
	res, err := h.Service.GetUser(ctx, user)
	if err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusInternalServerError, err))
		return
	}
	sendMessage(ctx, NewResult(res, http.StatusOK, nil))
}

func (h *Handler) GetSelfUser(ctx *gin.Context) {
	var user dto.DtoUuidUserFromWeb
	uuid := ctx.GetHeader("X-User-UUID")
	if uuid == "" {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, fmt.Errorf("invalid reqest")))
		return
	}
	user.UUID = uuid
	res, err := h.Service.GetUser(ctx, user)
	if err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusInternalServerError, err))
		return
	}
	sendMessage(ctx, NewResult(res, http.StatusOK, nil))
}

func (h *Handler) GetUsersList(ctx *gin.Context) {
	users, err := h.Service.GetUserList(ctx)
	if err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusInternalServerError, err))
		return
	}
	sendMessage(ctx, NewResult(users, http.StatusOK, nil))
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var user dto.DtoUserToDb
	uuid := ctx.GetHeader("X-User-UUID")
	if uuid == "" {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, fmt.Errorf("invalid reqest")))
		return
	}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, err))
		return
	}
	user.UUID = uuid
	err = h.Service.UpdateUser(ctx, user)
	if err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusInternalServerError, err))
		return
	}
	sendMessage(ctx, NewResult("Succes update", http.StatusOK, nil))
}

func (h *Handler) NotFound(ctx *gin.Context) {
	sendMessage(ctx, NewResult(nil, http.StatusNotFound, fmt.Errorf("not found")))
}
