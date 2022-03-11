package support

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env     string
	App     string
	Version string
	DB      string
}

func LoadConfig() *Config {
	config := Config{}
	err := config.loadEnvs()
	if err != nil {
		panic(err)
	}
	config.setEnvs()

	return &config
}

func (c *Config) loadEnvs() error {
	return godotenv.Load(".env")
}

func (c *Config) setEnvs() {
	c.Env = os.Getenv("ENVIRONMENT")
	c.App = os.Getenv("APP")
	c.Version = os.Getenv("VERSION")
	c.DB = os.Getenv("DB")
}
