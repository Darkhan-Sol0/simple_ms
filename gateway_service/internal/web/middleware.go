package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Timeout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if h.semaphore == nil {
			h.MakeSemophore()
		}
		ok, release := h.checkSemaphore(ctx)
		if !ok {
			ctx.AbortWithStatus(http.StatusRequestTimeout)
			return
		}
		defer release()
		ctx.Next()
	}
}

func (h *Handler) DecodeToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenAuth := ctx.GetHeader("Authorization")
		if tokenAuth == "" {
			ctx.Next()
			return
		}
		parts := strings.Split(tokenAuth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.Next()
			return
		}
		token := parts[1]
		jsonData, _ := json.Marshal(map[string]string{"token": token})
		link := fmt.Sprintf("%s/check_auth", h.Services.Service["auth"])
		req, _ := http.NewRequest("POST", link, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		cl := &http.Client{Timeout: time.Duration(h.Services.SemophoreTimeout) * time.Second}
		res, err := cl.Do(req)
		if err != nil {
			ctx.Next()
			return
		}
		defer res.Body.Close()
		resp, err := h.GetResponse(res)
		if err != nil {
			ctx.Next()
			return
		}
		if data, ok := resp.Data.(map[string]interface{}); ok {
			if uuid, ok := data["uuid"].(string); ok {
				ctx.Set("X-User-UUID", uuid)
			}
			if role, ok := data["user_role"].(string); ok {
				ctx.Set("X-User-Role", role)
			}
		}
		ctx.Next()
	}
}
