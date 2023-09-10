package entity

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ChatGptApiKey string `json:"chatGptApiKey"`
	UseLanguage   string `json:"UseLanguage"`
	CommitFormat  string `json:"CommitFormat"`
}

func ReadConfig() (*Config, error) {
	var result Config

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return &result, fmt.Errorf("error reading config file, %s", err.Error())
	}
	if err := viper.Unmarshal(&result); err != nil {
		return &result, fmt.Errorf("unable to decode into struct, %v", err.Error())
	}
	return &result, nil
}

func WriteConfig(config *Config) error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configMap := make(map[string]interface{})
	configBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshalling config: %s", err.Error())
	}
	err = json.Unmarshal(configBytes, &configMap)
	if err != nil {
		return fmt.Errorf("error unmarshalling config: %s", err.Error())
	}
	if err := viper.MergeConfigMap(configMap); err != nil {
		return err
	}
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("error saving config file, %s", err.Error())
	}
	return nil
}
