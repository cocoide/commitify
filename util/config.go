package util

import (
	"encoding/json"
	"fmt"

	"github.com/cocoide/commitify/internal/entity"
	"github.com/spf13/viper"
)

func ReadConfig() (*entity.Config, error) {
	var result entity.Config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return &result, fmt.Errorf("Error reading config file, %s", err.Error())
	}
	if err := viper.Unmarshal(&result); err != nil {
		return &result, fmt.Errorf("Unable to decode into struct, %v", err.Error())
	}
	return &result, nil
}

func WriteConfig(config *entity.Config) error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configMap := make(map[string]interface{})
	configBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("Error marshalling config: %s", err.Error())
	}
	err = json.Unmarshal(configBytes, &configMap)
	if err != nil {
		return fmt.Errorf("Error unmarshalling config: %s", err.Error())
	}
	if err := viper.MergeConfigMap(configMap); err != nil {
		return err
	}
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("Error saving config file, %s", err.Error())
	}
	return nil
}
