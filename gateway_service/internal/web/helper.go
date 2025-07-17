package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func GetResponce(ctx *gin.Context, response *http.Response) (Response, error) {
	defer response.Body.Close()
	var resp Response
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		sendMessage(ctx, ErrorResponce(errorDecodeJSON, http.StatusBadRequest, err))
		return Response{}, err
	}
	if resp.Err != nil {
		sendMessage(ctx, ErrorResponce(resp.Details, resp.Status, fmt.Errorf("error")))
		return Response{}, fmt.Errorf("error")
	}
	return resp, nil
}

func SendRequest(ctx *gin.Context, link, httpMethod string, jsonData []byte, headerMap map[string]string) (*http.Response, error) {
	validMethods := map[string]bool{
		"GET": true, "POST": true, "PUT": true,
		"PATCH": true, "DELETE": true,
	}
	if !validMethods[strings.ToUpper(httpMethod)] {
		return nil, fmt.Errorf("invalid HTTP method: %s", httpMethod)
	}
	req, err := http.NewRequestWithContext(ctx.Request.Context(), httpMethod, link, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	if jsonData != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for key, value := range headerMap {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	return client.Do(req)
}
