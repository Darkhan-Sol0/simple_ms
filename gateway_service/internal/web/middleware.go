package web

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

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
		token, err := extractToken(ctx)
		log.Println(token)
		if err != nil {
			ctx.Next()
			return
		}
		data, err := h.decodeToken(token)
		log.Println(data)
		if err != nil {
			ctx.Next()
			return
		}
		for key, value := range h.Services.ParsTags {
			if res, ok := data[key].(string); ok {
				ctx.Set(value, res)
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
		ctx.Set("link", link)
		ctx.Next()
	}
}
