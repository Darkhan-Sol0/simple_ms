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
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, err))
	}
	defer res.Body.Close()
	var resp Respones
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, err))
		return
	}
	if resp.Err != nil {
		sendMessage(ctx, NewResult(resp.Details, resp.Status, resp.Err))
		return
	}
	sendMessage(ctx, NewResult(resp.Data, resp.Status, nil))
}

func (h *Handler) Authorization(ctx *gin.Context) {

}
