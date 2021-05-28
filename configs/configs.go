package configs

import (
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	Api_Service     string   `json:"api-service"`
	User_Service    string   `json:"user-service"`
	Parking_Service string   `json:"parking_service"`
	Slot_Service    string   `json:"slot_service"`
	Booking_Service string   `json:"booking_service"`
	Payment_Service string   `json:"payment_service"`
	MQTT_Service    string   `json:"mqtt_service"`
	Database        Database `json:"db"`
	MQTT            MQTT
}

type Database struct {
	Db_Host string `json:"db_host"`
	Db_Name string `json:"db_name"`
}

type MQTT struct {
	Broker string `json:"broker"`
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
