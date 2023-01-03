package main

import (
	"fmt"
	"math/rand"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"github.com/qiangxue/fasthttp-routing"
)

type Map map[string]interface{}

func Generic(c *routing.Context, statusCode int, message string, data interface{}, success bool) {
	var code = -1
	if success {
		code = 0
	}

	r := IHTTPMessage{
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

func RandomToken(source interface{}) (string, error) {
	randNumber := rand.Intn(9999999 - 1000000) + 1000000
	hashed, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%v", source)), 10)
	if err != nil {
		return "", err
	}

	token := fmt.Sprintf("%d.%s", randNumber, hashed)

	return token, nil
}