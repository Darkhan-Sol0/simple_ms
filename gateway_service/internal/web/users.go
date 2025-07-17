package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"gateway/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserList(ctx *gin.Context) {
	link := fmt.Sprintf("%s/user_list", h.Services.Auth_service)
	response, err := SendRequest(ctx, link, GET, nil, nil)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorResponceInternalService, http.StatusBadRequest, err))
		return
	}
	resp, err := GetResponce(ctx, response)
	if err != nil {
		return
	}
	sendMessage(ctx, resp)
}

func (h *Handler) GetSelfUser(ctx *gin.Context) {
	userUUID, exists := ctx.Get("userUUID")
	if !exists {
		sendMessage(ctx, ErrorResponce(errorUUIDContext, http.StatusBadRequest, errors.New("user UUID not found in context")))
		return
	}
	header := map[string]string{"X-User-UUID": userUUID.(string)}
	link := fmt.Sprintf("%s/", h.Services.Users_service)
	response, err := SendRequest(ctx, link, GET, nil, header)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorResponceInternalService, http.StatusBadRequest, err))
		return
	}
	resp, err := GetResponce(ctx, response)
	if err != nil {
		return
	}
	sendMessage(ctx, resp)
}

func (h *Handler) UserUpdateInfo(ctx *gin.Context) {
	userUUID, exists := ctx.Get("userUUID")
	if !exists {
		sendMessage(ctx, ErrorResponce(errorUUIDContext, http.StatusBadRequest, errors.New("user UUID not found in context")))
		return
	}
	header := map[string]string{"X-User-UUID": userUUID.(string)}
	var userInfoUpdate dto.DtoUserUpdateToUser
	err := ctx.ShouldBindJSON(&userInfoUpdate)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorInvalidJSON, http.StatusBadRequest, err))
	}
	jsonData, err := json.Marshal(userInfoUpdate)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorMarshalJSON, http.StatusBadRequest, err))
		return
	}
	link := fmt.Sprintf("%s/", h.Services.Users_service)
	response, err := SendRequest(ctx, link, PATCH, jsonData, header)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorResponceInternalService, http.StatusBadRequest, err))
		return
	}
	resp, err := GetResponce(ctx, response)
	if err != nil {
		return
	}
	sendMessage(ctx, resp)
}
