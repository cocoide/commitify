package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setAPIKeyCmd = &cobra.Command{
	Use:   "set-apikey [api_key]",
	Short: "API Key settings for ChatGPT",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args [] string) {
		apikey := args[0]
		viper.Set("chatgpt.api_key", apikey)
		if err := viper.WriteConfig(); err != nil {
			log.Fatal("An error occurred while writing the configuration file:", err)
		}
		fmt.Println("ChatGPT API key has been set")
	},
}

var showAPIKeyCmd = &cobra.Command{
	Use: "show-apikey",
	Short: "Display ChatGPT API key",
	Run: func (cmd *cobra.Command, args [] string)  {
		apikey := viper.GetString("chatgpt.api_key")
		if apikey == "" {
			fmt.Println("API key is not set")
		} else {
			fmt.Println("ChatGPT APIKey:", apikey)
		}
	},
}

func init() {
	rootCmd.AddCommand(setAPIKeyCmd)
	rootCmd.AddCommand(showAPIKeyCmd)
}