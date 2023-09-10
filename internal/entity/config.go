package entity

import (
	"encoding/json"
	"fmt"

	pb "github.com/cocoide/commitify/pkg/grpc"
	"github.com/spf13/viper"
)

type Language int

const (
	JP Language = iota
	EN
)

type CodeFormat int

const (
	EmojiFormat CodeFormat = iota
	PrefixFormat
	NormalFormat
)

type AISource int

const (
	WrapServer AISource = iota
	OpenAiAPI
)

type Config struct {
	ChatGptApiKey string `json:"chatGpt_ApiKey"`
	UseLanguage   int    `json:"Use_Language"`
	CommitFormat  int    `json:"Commit_Format"`
	AISource      int    `json:"AI_Source"`
}

func (c *Config) Config2PbVars() (pb.CodeFormatType, pb.LanguageType) {
	var codeFormatType pb.CodeFormatType
	switch c.CommitFormat {
	case int(EmojiFormat):
		codeFormatType = pb.CodeFormatType_EMOJI
	case int(PrefixFormat):
		codeFormatType = pb.CodeFormatType_PREFIX
	default:
		codeFormatType = pb.CodeFormatType_NORMAL
	}

	var languageType pb.LanguageType
	switch c.UseLanguage {
	case int(JP):
		languageType = pb.LanguageType_JAPANESE
	default:
		languageType = pb.LanguageType_JAPANESE
	}

	return codeFormatType, languageType
}

func ReadConfig() (Config, error) {
	var result Config

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return result, fmt.Errorf("error reading config file, %s", err.Error())
	}
	if err := viper.Unmarshal(&result); err != nil {
		return result, fmt.Errorf("unable to decode into struct, %v", err.Error())
	}
	return result, nil
}

func WriteConfig(config Config) error {
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
