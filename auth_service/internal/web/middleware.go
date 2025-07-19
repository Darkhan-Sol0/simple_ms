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
