package main

const (
	Female = iota
	Male
	Other
)

type HTTPMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"exited_code"`
}