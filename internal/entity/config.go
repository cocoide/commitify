package entity

import (
	"encoding/json"
	"fmt"
	"os"

	pb "github.com/cocoide/commitify/proto/gen"
	"github.com/spf13/viper"
)

// コミットメッセージの言語の列挙型
type Language int

const (
	EN Language = iota
	JP
)

// コミットメッセージの形式の列挙型
type CodeFormat int

const (
	NormalFormat CodeFormat = iota
	EmojiFormat
	PrefixFormat
)

// ChatGPTのAPIを叩く場所
type GptRequestLocation int

const (
	Server GptRequestLocation = iota
	Client
	Qdrant
	Gemini
)

type Config struct {
	ChatGptApiKey string `json:"chatGptApiKey"`
	UseLanguage   int    `json:"UseLanguage"`
	CommitFormat  int    `json:"CommitFormat"`
	AISource      int    `json:"AISource"`
	GithubToken   string `json:"GithubToken"`
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
	homePath, err := os.UserHomeDir()
	if err != nil {
		return result, err
	}

	viper.AddConfigPath(homePath + "/.commitify")
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

func (c Config) WriteConfig() error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	viper.AddConfigPath(homePath + "/.commitify")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configMap := make(map[string]interface{})
	configBytes, err := json.Marshal(c)
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

func (c *Config) WithGithubToken(token string) *Config {
	c.GithubToken = token
	return c
}

func SaveConfig(configIndex, updateConfigParamInt int, updateConfigParamStr string) error {
	currentConfig, err := ReadConfig()
	if err != nil {
		return err
	}

	switch configIndex {
	case 0:
		currentConfig.ChatGptApiKey = updateConfigParamStr
	case 1:
		currentConfig.UseLanguage = updateConfigParamInt
	case 2:
		currentConfig.CommitFormat = updateConfigParamInt
	case 3:
		currentConfig.AISource = updateConfigParamInt
	}

	err = currentConfig.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) GptRequestLocation() GptRequestLocation {
	switch c.AISource {
	case 0:
		return Server
	case 1:
		return Client
	case 2:
		return Qdrant
	case 3:
		return Gemini
	default:
		return Server
	}
}
