package configuration

import (
	"flag"
	"os"
)

type Config struct {
	Address string `env:"RUN_ADDRESS"`
	DB      string `env:"DATABASE_URI"`
	Accural string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func NewConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Address, "a", "localhost:8000", "address for your server")
	flag.StringVar(&cfg.DB, "d", "", "connection for postgres")
	flag.StringVar(&cfg.Accural, "r", "localhost:9000", "accural system")

	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		cfg.Address = envRunAddr
	}

	if envDB := os.Getenv("DATABASE_URI"); envDB != "" {
		cfg.DB = envDB
	}

	if envAccural := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccural != "" {
		cfg.Accural = envAccural
	}
	return cfg
}
