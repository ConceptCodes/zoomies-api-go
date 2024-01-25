package config

import (
	"log"

	"github.com/spf13/viper"

	"zoomies-api-go/pkg/models"
)

type Config struct {
	Port          int               `mapstructure:"port"`
	SessionSecret string            `mapstructure:"session_secret"`
	Jwt           *models.JwtConfig `mapstructure:"jwt"`
	DB            *models.DBConfig  `mapstructure:"db"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
