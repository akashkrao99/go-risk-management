package config

type Config struct {
	Env string
	BlacklistedIps []string
	HttpServerConfig HttpServerConfig
}

type HttpServerConfig struct{
	Port string
}