package models

type Response struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	ErrorCode string      `json:"error_code"`
}
