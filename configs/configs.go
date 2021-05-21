package configs

import (
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	Api_Service  string `json:"api-service"`
	User_Service string `json:"user-service"`
}

var configuration *Configuration
var configError error

var configOnce sync.Once

func LoadConfigurations() (*Configuration, error) {
	configOnce.Do(func() {
		viper.SetConfigName("configs")
		viper.SetConfigType("json")
		viper.AddConfigPath("./configs")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			configError = err
		}

		err = viper.Unmarshal(&configuration)
		if err != nil {
			configError = err
		}
	})
	return configuration, configError
}
