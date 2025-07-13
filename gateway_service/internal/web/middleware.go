package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			sendMessage(ctx, NewResult("Unauthorized", http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		token := strings.Split(authHeader, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			sendMessage(ctx, NewResult("Invalid token", http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		jsonData, err := json.Marshal(token[1])
		if err != nil {
			sendMessage(ctx, NewResult("Failed to create JSON", http.StatusInternalServerError, err))
			ctx.Abort()
			return
		}
		res, err := http.Post(fmt.Sprintf("%s/check_auth", h.Services.Auth_service), "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			sendMessage(ctx, NewResult(nil, http.StatusUnauthorized, fmt.Errorf("token validation failed")))
			ctx.Abort()
			return
		}
		defer res.Body.Close()
		var resp Respones
		if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
			sendMessage(ctx, NewResult(nil, http.StatusBadRequest, err))
			ctx.Abort()
			return
		}
		if resp.Err != nil {
			sendMessage(ctx, NewResult(resp.Details, resp.Status, resp.Err))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
