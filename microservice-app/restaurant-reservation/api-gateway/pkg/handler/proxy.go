package handler

import (
	"io"
	"net/http"
	"reservation/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) ProxyHandler(ctx *gin.Context) {
	urlPath := ctx.Request.URL.Path
	requestMethod := ctx.Request.Method
	authBackendURL := viper.GetString("services.auth")
	paymentBackendURL := viper.GetString("services.payment")
	reservationBackendURL := viper.GetString("services.reservation")
	backendURL := ""

	switch {
	case strings.Contains(urlPath, "auth"):
		backendURL = authBackendURL + urlPath
	case strings.Contains(urlPath, "payment"):
		backendURL = paymentBackendURL + urlPath
	default:
		backendURL = reservationBackendURL + urlPath
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
