package server

import (
	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello, World!")
}