package main

import (
	"github.com/mehulgohil/go-bffauth.git/authenticator"
	"github.com/mehulgohil/go-bffauth.git/config"
	"log"
)

func main() {
	//initialize env variables
	config.LoadEnvVariables()

	//initialize redis
	config.Redis().InitRedisConnection()

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	//initialize api routes
	app := Router().InitRouter(auth, config.Redis().(*config.RedisHandler).RedisClient)

	err = app.Listen(":" + config.EnvVariables.AppPort)
	if err != nil {
		panic("unable to start server")
	}
}
