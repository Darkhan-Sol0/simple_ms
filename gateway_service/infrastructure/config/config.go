package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Services struct {
	Auth_service  string `env:"AUTH_SERVICE_URL"`
	Users_service string `env:"USERS_SERVICE_URL"`
}

var (
	services *Services
	once     sync.Once
)

func parsServiceConf() *Services {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: error loading .env file: ", err)
		return nil
	}
	services = &Services{}
	if err := cleanenv.ReadConfig(".env", services); err != nil {
		log.Println("Warning: error reading Services config: ", err)
		cleanenv.GetDescription(services, nil)
	}
	return services
}

func GetService() *Services {
	once.Do(func() {
		parsServiceConf()
	})
	return services
}
