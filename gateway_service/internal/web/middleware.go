package web

import (
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
			sendMessage(ctx, ErrorResponce(errorUnauthorized, http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		token := strings.Split(authHeader, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			sendMessage(ctx, ErrorResponce(errorToken, http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		jsonData, err := json.Marshal(token[1])
		if err != nil {
			sendMessage(ctx, ErrorResponce(errorMarshalJSON, http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		link := fmt.Sprintf("%s/check_auth", h.Services.Auth_service)
		response, err := SendRequest(ctx, link, POST, jsonData, nil)
		if err != nil {
			sendMessage(ctx, ErrorResponce(errorToken, http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		resp, err := GetResponce(ctx, response)
		if err != nil {
			return
		}
		var userData dto.DtoUserAuthToken
		if data, ok := resp.Data.(map[string]interface{}); ok {
			userData = dto.DtoUserAuthToken{
				UUID: data["uuid"].(string),
				Role: data["user_role"].(string),
			}
		} else {
			sendMessage(ctx, ErrorResponce(errorToken, http.StatusUnauthorized, fmt.Errorf("unauthorized")))
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
			sendMessage(ctx, ErrorResponce(errorRoleNotFound, http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		userRoleStr, ok := userRole.(string)
		if !ok {
			sendMessage(ctx, ErrorResponce(errorRoleType, http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		if userRoleStr != role {
			sendMessage(ctx, ErrorResponce(errorAccessRole, http.StatusUnauthorized, fmt.Errorf("unauthorized")))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
