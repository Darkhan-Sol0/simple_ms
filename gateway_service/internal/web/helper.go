package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) proxyRequest(ctx *gin.Context, link string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx.Request.Context(), ctx.Request.Method, link, ctx.Request.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	for _, header := range h.Services.ParsTags {
		if value, ok := ctx.Get(header); ok {
			req.Header.Add(header, value.(string))
		}
	}
	client := &http.Client{
		Timeout: time.Duration(h.Services.RequestTimeout) * time.Second,
	}
	return client.Do(req)
}

func (h *Handler) proxyResponse(ctx *gin.Context, response Response) {
	if response.Err != nil {
		h.sendMessage(ctx, NewResult(nil, response.Status, fmt.Errorf("error: %s", response.Err)))
	}
	h.sendMessage(ctx, NewResult(response.Data, response.Status, nil))
}

func (h *Handler) GetResponse(response *http.Response) (Response, error) {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{Status: response.StatusCode}, fmt.Errorf("error: %s", err)
	}
	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		return Response{Status: response.StatusCode}, fmt.Errorf("error: %s", err)
	}
	return res, nil
}

func (h *Handler) checkSemaphore(c *gin.Context) (bool, func()) {
	select {
	case h.semaphore <- struct{}{}:
		return true, func() { <-h.semaphore }
	case <-time.After(time.Duration(h.Services.SemophoreTimeout) * time.Second):
		return false, nil
	case <-c.Request.Context().Done():
		return false, nil
	}
}

func extractToken(ctx *gin.Context) (string, error) {
	tokenAuth := ctx.GetHeader("Authorization")
	if tokenAuth == "" {
		return "", fmt.Errorf("error: No token")
	}
	parts := strings.Split(tokenAuth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("error: Incorrect token")
	}
	token := parts[1]
	return token, nil
}

func (h *Handler) decodeToken(token string) (map[string]interface{}, error) {
	jsonData, _ := json.Marshal(map[string]string{"token": token})
	req, _ := http.NewRequest("POST", h.Services.URLTokenDecoder, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Internal-Secret", h.Services.InternalTag)
	cl := &http.Client{Timeout: time.Duration(h.Services.SemophoreTimeout) * time.Second}
	res, err := cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}
	defer res.Body.Close()
	resp, err := h.GetResponse(res)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}
	if data, ok := resp.Data.(map[string]interface{}); ok {
		return data, nil
	}
	return nil, fmt.Errorf("error: No token")
}
