package web

import (
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
		sendMessage(ctx, NewResult(res, http.StatusBadRequest, err))
	}
	sendMessage(ctx, NewResult(res.Body, http.StatusCreated, nil))
}
