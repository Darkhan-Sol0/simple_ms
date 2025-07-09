package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MainHandler(ctx *gin.Context) {
	text := "Main Text"
	sendMessage(ctx, NewResult(text, http.StatusOK, nil))
}

func (h *Handler) Test(ctx *gin.Context) {
	res, err := http.Get(fmt.Sprintf("%s/", h.Services.Auth_service))
	if err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, err))
	}
	var dec struct {
		Data   any `json:"data"`
		Status int `json:"status"`
	}
	if err := json.NewDecoder(res.Body).Decode(&dec); err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusInternalServerError, err))
		return
	}
	sendMessage(ctx, NewResult(dec, http.StatusOK, nil))
}
