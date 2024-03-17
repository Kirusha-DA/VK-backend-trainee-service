package config

import (
	"sync"

	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool         `yaml:"is_debug" env-required:"true"`
	Listen  ListenConfig  `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type ListenConfig struct {
	Type   string `yaml:"type" env-default:"tcp"`
	BindIp string `yaml:"bind_ip" env-default:"127.0.0.1"`
	Port   string `yaml:"port" env-default:"8080"`
}

type StorageConfig struct {
	Host     string `yaml:"host" env-default:"5432"`
	Port     string `yaml:"port" env-default:"5432"`
	DB       string `yaml:"database" env-default:"filmoteka"`
	User     string `yaml:"username" env-default:"postgres"`
	Password string `yaml:"password" env-default:"123"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read app config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("configs/config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
