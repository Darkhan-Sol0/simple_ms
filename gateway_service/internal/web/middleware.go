package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
		link := fmt.Sprintf("%s/internal/check_auth", h.Services.Service["auth"])
		req, _ := http.NewRequest("POST", link, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Internal-Secret", h.Services.InternalTag)
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
			for key, value := range h.Services.ParsTags {
				if res, ok := data[key].(string); ok {
					ctx.Set(value, res)
				}
			}
		}
		ctx.Next()
	}
}

func (h *Handler) CheckService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service := ctx.Param("service")
		serviceURL, exists := h.Services.Service[service]
		if !exists {
			h.sendMessage(ctx, NewResult(nil, http.StatusNotFound, fmt.Errorf("error: Not Found Service")))
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		path := strings.TrimPrefix(ctx.Param("path"), "/")
		link, err := url.JoinPath(serviceURL, path)
		if err != nil {
			h.sendMessage(ctx, NewResult(nil, http.StatusNotFound, fmt.Errorf("error: Not Found Service")))
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Println(link)
		ctx.Set("link", link)
		ctx.Next()
	}
}
