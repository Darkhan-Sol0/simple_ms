package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Registration(ctx *gin.Context) {
	var regUser dto.DtoRegUserLogin
	if err := ctx.ShouldBindJSON(&regUser); err != nil {
		sendMessage(ctx, NewResult("invalid request body", http.StatusBadRequest, err))
		return
	}
	jsonData, _ := json.Marshal(regUser)
	res, err := http.Post(fmt.Sprintf("%s/sign_up", h.Services.Auth_service), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		sendMessage(ctx, NewResult(res, http.StatusBadRequest, err))
	}
	sendMessage(ctx, NewResult(res, http.StatusCreated, nil))
}

func (h *Handler) Authorization(ctx *gin.Context) {

}
