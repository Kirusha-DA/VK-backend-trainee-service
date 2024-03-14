package config

import (
	"sync"

	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/spf13/viper"
)

type Config struct {
	IsDebug bool
	Listen  struct {
		Type   string
		BindIp string
		Port   string
	}
	Storage StorageConfig
}

type StorageConfig struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read app config")
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal(err)
		}
		instance = &Config{
			IsDebug: viper.GetBool("is_debug"),
			Listen: struct {
				Type   string
				BindIp string
				Port   string
			}{viper.GetString("listen.type"), viper.GetString("listen.bind_ip"), viper.GetString("listen.port")},
		}
	})

	return instance
}
