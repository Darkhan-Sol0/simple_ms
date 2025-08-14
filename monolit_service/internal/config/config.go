package config

import (
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	once   sync.Once
	config Config
)

type (
	ServerConfig struct {
		ServerHTTP `yaml:"server_http"`
		DatabasePG `yaml:"database"`
	}
	ServerHTTP struct {
		Address        string        `yaml:"address"`
		SessionTimeout time.Duration `yaml:"session_timeout"`
		IdleTimeout    time.Duration `yaml:"idle_timeout"`
	}

	DatabasePG struct {
		Env      string `yaml:"database_env"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database_name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	Config interface {
		GetAddress() string
		GetSessionTimeout() time.Duration
		GetIdleTimeout() time.Duration

		GetDBEnv() string
		GetDBPort() string
		GetDBHost() string
		GetDBDatabase() string
		GetDBUsername() string
		GetDBPassword() string
	}
)

func LoadConfig() Config {
	once.Do(func() {
		config = &ServerConfig{}
		configPath := "config/config.yaml"
		if err := cleanenv.ReadConfig(configPath, config); err != nil {
			log.Fatalf("error read config %s: %v", configPath, err)
		}
	})
	return config
}

func (s *ServerConfig) GetAddress() string {
	return s.ServerHTTP.Address
}

func (s *ServerConfig) GetSessionTimeout() time.Duration {
	return s.ServerHTTP.SessionTimeout
}

func (s *ServerConfig) GetIdleTimeout() time.Duration {
	return s.ServerHTTP.IdleTimeout
}

func (s *ServerConfig) GetDBEnv() string {
	return s.DatabasePG.Env
}

func (s *ServerConfig) GetDBPort() string {
	return s.DatabasePG.Port
}

func (s *ServerConfig) GetDBHost() string {
	return s.DatabasePG.Host
}

func (s *ServerConfig) GetDBDatabase() string {
	return s.DatabasePG.Database
}

func (s *ServerConfig) GetDBUsername() string {
	return s.DatabasePG.Username
}

func (s *ServerConfig) GetDBPassword() string {
	return s.DatabasePG.Password
}
