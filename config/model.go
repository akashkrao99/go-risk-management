package config

type Config struct {
	Env string
	HttpServerConfig HttpServerConfig
}

type HttpServerConfig struct{
	Port string
}