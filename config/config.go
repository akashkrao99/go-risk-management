package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once
var config *Config

func InitializeConfig() {
	once.Do(func() {

		if config != nil {
			return
		}
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		config = &Config{
			Env: os.Getenv("ENV"),
			HttpServerConfig: HttpServerConfig{
				Port: os.Getenv("HTTP_SERVER_PORT"),
			},
		}
	})

}

// following singleton pattern
func GetConfig() *Config {
	if config == nil {
		InitializeConfig()
	}
	return config
}
