package main

import (
	"encoding/json"

	"github.com/qiangxue/fasthttp-routing"
)

func Generic(c *routing.Context, statusCode int, message string, data interface{}, success bool) {
	var code = -1
	if success {
		code = 0
	}

	r := HTTPMessage{
		Message: message,
		Data:    data,
		Code:    code,
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		c.SetStatusCode(500)
		return
	}

	c.SetStatusCode(statusCode)
	c.Write(bytes)
}