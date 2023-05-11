package config

import "github.com/caarlos0/env/v8"

type envConfig struct {
	AppPort           string `env:"APP_PORT"`
	Auth0Domain       string `env:"AUTH0_DOMAIN"`
	Auth0ClientID     string `env:"AUTH0_CLIENT_ID"`
	Auth0ClientSecret string `env:"AUTH0_CLIENT_SECRET"`
	Auth0CallbackURL  string `env:"AUTH0_CALLBACK_URL"`
	Auth0Audience     string `env:"AUTH0_AUDIENCE"`
	BackendApi        string `env:"BACKEND_API"`
	FrontendURL       string `env:"FRONTEND_URL"`
	RedisHost         string `env:"REDIS_HOST"`
	RedisPassword     string `env:"REDIS_PASSWORD"`
}

var EnvVariables envConfig

func LoadEnvVariables() {
	if err := env.Parse(&EnvVariables); err != nil {
		panic("unable to load environment variables")
	}
}
