package main

import (
	"log"
	"os"
	"fmt"
	"runtime"

	"github.com/joho/godotenv"
)

var hash_key string

type Core struct {
	production bool
	logger *log.Logger
	db *Database
	redis *Redis
	server *HTTPServer
	hash_key *string
}

func (c *Core) Run() {
	c.LoadEnviromentVariables()
	c.logger = log.New(os.Stdout, "[CHATME-API] ", 0)

	db := Database{}
	db.Connect(c)

	redis := Redis{}
	redis.Init(c)

	httpServer := HTTPServer{
		addr: "0.0.0.0",
		port: os.Getenv("PORT"),
	}

	httpServer.RunHTTPServer(c)
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

	hash_key, err = RandomToken(fmt.Sprintf("%d:%d:%s", os.Getpid(), runtime.NumCPU(), runtime.GOROOT()))
	if err != nil {
		panic(err)
	}

	os.Setenv("HASH_KEY", hash_key)
	c.hash_key = &hash_key
}