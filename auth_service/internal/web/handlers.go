package web

import (
	"auth_service/internal/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Registration(ctx *gin.Context) {
	var regUser dto.DtoRegUserFromWeb
	if err := ctx.ShouldBindJSON(&regUser); err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error())))
		return
	}
	uuid, err := h.Service.CreateUser(ctx, regUser)
	if err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, fmt.Errorf("invalid create user: %s", err.Error())))
		return
	}
	jsonData, _ := json.Marshal(uuid)
	res, err := h.proxyRequest(ctx, "http://user_service:8282/", "POST", io.NopCloser(bytes.NewBuffer(jsonData)))
	if err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, fmt.Errorf("err: %s", err.Error())))
		return
	}
	data, err := h.GetResponse(res)
	if err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, fmt.Errorf("err: %s", err.Error())))
		return
	}
	sendMessage(ctx, NewResult(data, http.StatusCreated, nil))
}

func (h *Handler) Authorization(ctx *gin.Context) {
	var authUser dto.DtoAuthUser
	if err := ctx.ShouldBindJSON(&authUser); err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error())))
		return
	}
	token, err := h.Service.AuthUser(ctx, authUser)
	if err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, fmt.Errorf("invalid auth user: %s", err.Error())))
		return
	}
	sendMessage(ctx, NewResult(token, http.StatusCreated, nil))
}

func (h *Handler) CheckAuthorization(ctx *gin.Context) {
	var token dto.DtoTokenChecker
	if err := ctx.ShouldBindJSON(&token); err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error())))
		return
	}
	userOut, err := h.Service.TokenChecker(ctx, token.Token)
	if err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusUnauthorized, fmt.Errorf("invalid request body: %s", err.Error())))
		return
	}
	sendMessage(ctx, NewResult(userOut, http.StatusOK, nil))
}

func (h *Handler) GetUsersList(ctx *gin.Context) {
	userList, err := h.Service.GetUsersList(ctx)
	if err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, fmt.Errorf("invalid internal service: %s", err.Error())))
		return
	}
	sendMessage(ctx, NewResult(userList, http.StatusOK, nil))
}

func (h *Handler) NotFound(ctx *gin.Context) {
	sendMessage(ctx, NewResult(nil, http.StatusNotFound, fmt.Errorf("not found")))
}
