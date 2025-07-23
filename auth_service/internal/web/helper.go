package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data   any `json:"data"`
	Err    any `json:"error"`
	Status int
}

func (h *Handler) proxyRequest(ctx *gin.Context, link, method string, body io.ReadCloser) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx.Request.Context(), method, link, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("X-Internal-Secret", "BOOBIES")
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	return client.Do(req)
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
