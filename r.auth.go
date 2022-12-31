package main

import (
	"github.com/qiangxue/fasthttp-routing"
)

type AuthRouter struct {
	core *Core
	router *routing.RouteGroup
} 

func (r *AuthRouter) SetRoutes() {
	r.SetEndpoints()
}

func (r *AuthRouter) SetEndpoints() {
	r.router.Get("/session", r.GetSession) // Get session information
	r.router.Post("/session", nil) // Create new session
	r.router.Delete("/session", nil) // Delete session
}

func (r *AuthRouter) GetSession(c *routing.Context) error {
	c.Write([]byte("<session>"))

	return nil
}

func (r *AuthRouter) CreateSession(c *routing.Context) error {
	return nil
}