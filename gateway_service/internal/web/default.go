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
	res, err := http.Get(fmt.Sprintf("%s/suc", h.Services.Auth_service))
	if err != nil {
		sendMessage(ctx, NewResult(err.Error(), http.StatusBadRequest, err))
	}
	defer res.Body.Close()
	var resp Response
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		sendMessage(ctx, NewResult(resp.Data, resp.Status, err))
		return
	}
	if resp.Err != nil {
		sendMessage(ctx, NewResult(resp.Details, resp.Status, resp.Err))
		return
	}
	sendMessage(ctx, NewResult(resp.Data, resp.Status, nil))
}

func (h *Handler) Test_bad(ctx *gin.Context) {
	res, err := http.Get(fmt.Sprintf("%s/err", h.Services.Auth_service))
	if err != nil {
		sendMessage(ctx, NewResult(nil, http.StatusBadRequest, err))
	}
	defer res.Body.Close()
	var resp Response
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
