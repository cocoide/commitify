package main

import (
	"log"

	"github.com/cocoide/commitify/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("An error occurred while reading the configuration file:", err)
	}

	cmd.Execute()
}
