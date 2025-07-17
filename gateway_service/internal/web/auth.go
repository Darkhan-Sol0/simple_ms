package web

import (
	"encoding/json"
	"fmt"
	"gateway/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Registration(ctx *gin.Context) {
	var regUser dto.DtoRegUserLogin
	if err := ctx.ShouldBindJSON(&regUser); err != nil {
		sendMessage(ctx, ErrorResponce(errorInvalidJSON, http.StatusBadRequest, err))
		return
	}
	jsonData, err := json.Marshal(regUser)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorMarshalJSON, http.StatusBadRequest, err))
		return
	}
	link := fmt.Sprintf("%s/sign_up", h.Services.Auth_service)
	responseAuth, err := SendRequest(ctx, link, POST, jsonData, nil)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorResponceInternalService, http.StatusBadRequest, err))
	}
	resp, err := GetResponce(ctx, responseAuth)
	if err != nil {
		return
	}
	jsonData, err = json.Marshal(resp.Data)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorMarshalJSON, http.StatusBadRequest, err))
		return
	}
	link = fmt.Sprintf("%s/", h.Services.Users_service)
	responseUser, err := SendRequest(ctx, link, POST, jsonData, nil)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorResponceInternalService, http.StatusBadRequest, err))
		return
	}
	resp, err = GetResponce(ctx, responseUser)
	if err != nil {
		return
	}
	sendMessage(ctx, resp)
}

func (h *Handler) Authorization(ctx *gin.Context) {
	var regUser dto.DtoAuthUser
	if err := ctx.ShouldBindJSON(&regUser); err != nil {
		sendMessage(ctx, ErrorResponce(errorInvalidJSON, http.StatusBadRequest, err))
		return
	}
	jsonData, err := json.Marshal(regUser)
	if err != nil {
		sendMessage(ctx, ErrorResponce(errorMarshalJSON, http.StatusBadRequest, err))
		return
	}
	link := fmt.Sprintf("%s/sign_in", h.Services.Auth_service)
	response, err := SendRequest(ctx, link, POST, jsonData, nil)
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
