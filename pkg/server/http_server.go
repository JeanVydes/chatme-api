package server

import (
	"github.com/valyala/fasthttp"
)

func run() {
	fasthttp.ListenAndServe(":8080", requestHandler)
}