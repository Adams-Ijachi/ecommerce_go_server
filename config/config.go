package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DatabaseURL struct {
		Url string
	}
}

func Init() *Configuration {
	var envFile string

	flag.StringVar(&envFile, "env", "", "Env Variable File")
	flag.Parse()

	if envFile != "" {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	return &Configuration{}
}

func (c *Configuration) LoadEnvVariables() error {

	c.DatabaseURL.Url = os.Getenv("DATABASE_URL")

	return nil

}
