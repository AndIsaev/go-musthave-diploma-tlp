package configuration

import (
	"flag"
	"os"
)

type Config struct {
	Address string `env:"RUN_ADDRESS"`
	DB      string `env:"DATABASE_URI"`
	Accrual string `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func NewConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Address, "a", "localhost:8000", "address for your server")
	flag.StringVar(&cfg.DB, "d", "", "connection for postgres")
	flag.StringVar(&cfg.Accrual, "r", "localhost:8080", "accrual system")

	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		cfg.Address = envRunAddr
	}

	if envDB := os.Getenv("DATABASE_URI"); envDB != "" {
		cfg.DB = envDB
	}

	if envAccrual := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccrual != "" {
		cfg.Accrual = envAccrual
	}
	return cfg
}
