package handler

import (
	"farmish/pkg/response"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ProxyHandler(ctx *gin.Context) {
	urlPath := ctx.Request.URL.Path
	requestMethod := ctx.Request.Method
	authBackendURL := "http://localhost:8081"
	farmBackendURL := "http://localhost:8082"
	backendURL := ""

	switch {
	case strings.Contains(urlPath, "auth"):
		backendURL = authBackendURL + urlPath
	default:
		backendURL = farmBackendURL + urlPath
	}

	req, err := http.NewRequest(requestMethod, backendURL, ctx.Request.Body)

	if err != nil {
		response.NewErrorResponse(ctx, req.Response.StatusCode, "error creating request")
		return
	}

	for key, value := range ctx.Request.Header {
		req.Header[key] = value
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadGateway, "error contacting backend")
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	ctx.Data(res.StatusCode, res.Header.Get("Content-Type"), body)
}
