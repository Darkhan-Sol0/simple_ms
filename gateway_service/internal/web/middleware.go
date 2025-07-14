package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/internal/dto"
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
		var response Response
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			sendMessage(ctx, NewResult(nil, http.StatusBadRequest, err))
			ctx.Abort()
			return
		}
		if response.Err != nil {
			sendMessage(ctx, NewResult(response.Details, response.Status, response.Err))
			ctx.Abort()
			return
		}
		var userData dto.DtoUserAuthToken
		if data, ok := response.Data.(map[string]interface{}); ok {
			userData = dto.DtoUserAuthToken{
				UUID: data["uuid"].(string),
				Role: data["role"].(string),
			}
		} else {
			sendMessage(ctx, NewResult("bad conver token", http.StatusBadRequest, fmt.Errorf("bad conver token")))
			ctx.Abort()
			return
		}
		ctx.Set("userUUID", userData.UUID)
		ctx.Set("userRole", userData.Role)
		ctx.Next()
	}
}

func (h *Handler) RoleAccessor(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("userRole")
		if !exists {
			sendMessage(ctx, NewResult("user role not found", http.StatusUnauthorized, fmt.Errorf("user role not found")))
			ctx.Abort()
			return
		}
		userRoleStr, ok := userRole.(string)
		if !ok {
			sendMessage(ctx, NewResult("invalid user role type", http.StatusInternalServerError, fmt.Errorf("invalid user role type")))
			ctx.Abort()
			return
		}
		if userRoleStr != role {
			sendMessage(ctx, NewResult("access denied", http.StatusForbidden, fmt.Errorf("access denied")))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
