package configuration

import (
	"flag"
	"os"
)

type Config struct {
	Address string `env:"RUN_ADDRESS"`
}

func NewConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Address, "a", "localhost:8000", "address for your server")

	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		cfg.Address = envRunAddr
	}
	return cfg
}
