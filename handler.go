package main

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type IV1RouterGroup struct {
	Current *routing.RouteGroup
	Users *routing.RouteGroup
}

type IRouterAPIGroup struct {
	Current *routing.RouteGroup
	Auth *routing.RouteGroup

	// Sub Collections
	V1 *IV1RouterGroup
}

type IRouterGroups struct {
	API IRouterAPIGroup
}

type Router struct {
	core *Core
	router *routing.Router
	groups IRouterGroups
}

func (r *Router) New() func(c *fasthttp.RequestCtx) {
	r.router = routing.New()
	r.SetGroups()
	r.SetRoutes(r.core)

	return r.router.HandleRequest
}

func (r *Router) SetGroups() {
	api := r.router.Group("/api")
	v1 := api.Group("/v1")

	r.groups = IRouterGroups{
		API: IRouterAPIGroup{
			Current: api,
			V1: &IV1RouterGroup{
				Current: v1,
				Users: v1.Group("/users"),
			},
			Auth: api.Group("/auth"),
		},
	}
}

func (r *Router) SetRoutes(core *Core) {
	auth := AuthRouter{
		core: core,
		router: r.groups.API.Auth,
	}

	users := UsersRouter{
		core: core,
		router: r.groups.API.V1.Users,
	}

	auth.SetRoutes()
	users.SetRoutes()
}