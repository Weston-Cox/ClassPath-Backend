//* internal/config/config.go:
//****************************************************************************************
//* Contains configuration loading logic,
//* such as reading environment variables or configuration files.
//****************************************************************************************

package config

import (
	"os"
)

type Config struct {
    ServerAddress string
    DatabaseURL   string
}

func LoadConfig() (*Config, error) {
    return &Config{
        ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
        DatabaseURL:   getEnv("DATABASE_URL", "defaulturl"),
    }, nil
}

func getEnv(key, defaultValue string) (string){
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
