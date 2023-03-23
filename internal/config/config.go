package config

import (
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBHost      string        `env:"POSTGRES_HOST" env-default:"localhost"`
	DBPort      string        `env:"POSTGRES_PORT" env-default:"5432"`
	DBUser      string        `env:"POSTGRES_USER"`
	DBPass      string        `env:"POSTGRES_PASS"`
	DBName      string        `env:"POSTGRES_NAME"`
	IP          string        `env:"IP"`
	Port        string        `env:"PORT"`
	Limit       int           `env:"LIMIT"`
	Cooldown    time.Duration `env:"COOLDOWN"`
	NetPrefix   int           `env:"NETPREFIX"`
	CacheExpire time.Duration `env:"CACHEEXPIRE"`
}

var instance *Config
var once sync.Once

// func GetConfig() *Config {
// 	once.Do(func() {
// 		cfg := Config{}
// 		if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
// 			log.Fatal("failed to read config", err)
// 		}
// 		instance = &cfg
// 	})
// 	if instance == nil {
// 		log.Fatal("failed to initialize config")
// 	}
// 	return instance
// }

func GetConfig() *Config {
	once.Do(func() {
		cfg := Config{}
		if err := cleanenv.ReadConfig("../.env", &cfg); err != nil {
			log.Fatal("failed to read config", err)
		}
		instance = &cfg
	})
	if instance == nil {
		log.Fatal("failed to initialize config")
	}
	return instance
}
