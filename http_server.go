package main

import (
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

type HTTPServer struct {
	service *fasthttp.Server
	addr string
	port string
}

func (h *HTTPServer) RunHTTPServer(core *Core) {
	h.service = &fasthttp.Server{
		Handler: requestHandler,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
		LogAllErrors: true,
		Logger: log.New(nil, "[CHATME-API]", 0),
	}

	core.server = h
	h.service.ListenAndServe(fmt.Sprintf("%s:%s", h.addr, h.port))
}