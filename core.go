package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Core struct {
	production bool
}

func (c *Core) Run() {
	c.LoadEnviromentVariables()

	httpServer := HTTPServer{
		addr: "0.0.0.0",
		port: os.Getenv("PORT"),
	}

	httpServer.RunHTTPServer()
}

func (c *Core) LoadEnviromentVariables() {
	envFilePath := ".env.development"
	if c.production {
		envFilePath = ".env.production"
	}

	err := godotenv.Load(envFilePath)	
	if err != nil {
		panic(err)
	}
}