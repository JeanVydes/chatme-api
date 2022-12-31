package main

import (
	"os"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func (r *Redis) Init(c *Core) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	r.client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB: 0,
	})

	_, err := r.client.Ping(r.client.Context()).Result()
	if err != nil {
		panic(err)
	}
	
	c.redis = r
	c.logger.Println("REDIS CONNECTED AT " + host + ":" + port)
} 