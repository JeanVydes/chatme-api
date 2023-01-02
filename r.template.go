package main

import (
	"github.com/qiangxue/fasthttp-routing"
)

type ITemplateRouter struct {
	//core *Core
	router *routing.RouteGroup
} 

func (r *ITemplateRouter) SetRoutes() {
	r.SetEndpoints()
}

func (r *ITemplateRouter) SetEndpoints() {
	r.router.Get("/", nil)
	r.router.Post("/", nil)
	r.router.Delete("/", nil)
}
