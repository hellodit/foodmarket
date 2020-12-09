package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Read config from yaml file
func Read() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

}
