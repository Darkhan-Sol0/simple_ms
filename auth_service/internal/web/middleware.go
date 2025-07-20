package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoleChecker(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleHeader := ctx.Request.Header.Get("X-User-Role")
		if roleHeader != role {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}

func (h *Handler) Internal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tag := ctx.GetHeader("X-Internal-Secret")
		if tag != "BOOBIES" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
