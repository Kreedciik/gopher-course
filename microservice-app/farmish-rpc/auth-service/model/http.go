package model

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPSuccess struct {
	Message string `json:"message"`
}

type HTTPDataSuccess struct {
	Data interface{} `json:"data"`
}
