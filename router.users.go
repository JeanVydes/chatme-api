package main

import (
	"github.com/qiangxue/fasthttp-routing"
)

type UsersRouter struct {
	core *Core
	router *routing.RouteGroup
} 

func (r *UsersRouter) SetRoutes() {
	r.SetEndpoints()
}

func (r *UsersRouter) SetEndpoints() {
	r.router.Get("", AuthMiddleware(r.GetUser, false)) // Get user
	r.router.Post("", r.CreateUser) // Create new user record in SQL database
}

func (r *UsersRouter) GetUser(c *routing.Context) error {
	result, err := r.ProcessGetUser(c)
	if err != nil {
		Generic(c, 500, err.Error(), nil, false)
		return nil
	}

	Generic(c, 200, "OK", result, true)

	return nil
}

func (r *UsersRouter) CreateUser(c *routing.Context) error {
	result, err := r.ProcessCreateUser(c)
	if err != nil {
		Generic(c, 500, err.Error(), nil, false)
		return nil
	}

	Generic(c, 200, "OK", result, true)

	return nil
}