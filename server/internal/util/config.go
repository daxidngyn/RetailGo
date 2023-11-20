package util

import (
	"fmt"
	"os"
)

type Config struct {
	DBDriver      string // `"DB_DRIVER"`
	DBSource      string // "DB_SOURCE"`
	ServerAddress string // `"SERVER_ADDRESS"`
	ClerkSK       string // `"CLERK_SK"`
	RedisAddress  string // `"REDIS_ADDRESS"`
	StripeSK      string // `"STRIPE_SK"`
}

func LoadConfig() (config Config, err error) {
	// viper.AddConfigPath("./util/config")
	// viper.SetConfigName(".env")
	// viper.SetConfigType("env")

	// viper.AutomaticEnv()

	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBSource = os.Getenv("DB_SOURCE")
	config.ServerAddress = envPortOr("8080")
	config.ClerkSK = os.Getenv("CLERK_SK")
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	config.StripeSK = os.Getenv("STRIPE_SK")

	fmt.Println(config)

	return
}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return port
}
