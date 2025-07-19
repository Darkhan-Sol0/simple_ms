package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) proxyRequest(ctx *gin.Context, link string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx.Request.Context(), ctx.Request.Method, link, ctx.Request.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	headersToCopy := []string{
		"X-User-UUID",
		"X-User-Role",
	}

	for _, header := range headersToCopy {
		if value, ok := ctx.Get(header); ok {
			req.Header.Add(header, value.(string))
		}
	}
	client := &http.Client{
		Timeout: time.Duration(h.Services.RequestTimeout) * time.Second,
	}
	return client.Do(req)
}

func (h *Handler) proxyResponse(ctx *gin.Context, response *http.Response) {
	h.sendMessage(ctx, response)
}

func (h *Handler) GetResponse(response *http.Response) (Response, error) {
	defer response.Body.Close()
	var resp Response
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return Response{Status: http.StatusBadGateway}, err
	}
	if resp.Err != nil {
		return Response{Status: resp.Status}, fmt.Errorf("error: %s", resp.Details)
	}
	return resp, nil
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
