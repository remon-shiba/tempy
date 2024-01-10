package model

var AnyResponse map[string]interface{}

type (
	ResponseBody struct {
		Status    int         `json:"status"`
		Message   string      `json:"message"`
		IsSuccess bool        `json:"isSuccess,omitempty"`
		Error     error       `json:"Error,omitempty"`
		Data      interface{} `json:"data,omitempty"`
		Other     interface{} `json:"other,omitempty"`
	}
)
