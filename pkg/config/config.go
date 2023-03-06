package config

import (
	"github.com/spf13/viper"
	"log"
)

func FromEnv(token string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(token).(string)
	if !ok {
		log.Fatal(err)
	}

	return value
}
