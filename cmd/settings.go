package cmd

import (
	"fmt"
	"log"

	"github.com/cocoide/commitify/util"
	"github.com/spf13/cobra"
)

var setAPIKeyCmd = &cobra.Command{
	Use:   "set-apikey [api_key]",
	Short: "API Key settings for ChatGPT",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config, _ := util.ReadConfig()
		config.ChatGptApiKey = args[0]
		if err := util.WriteConfig(config); err != nil {
			log.Fatal("Failed to write into config", err)
		}
		fmt.Println("ChatGPT Token has been set")
	},
}

var showAPIKeyCmd = &cobra.Command{
	Use:   "show-apikey",
	Short: "Display ChatGPT API key",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := util.ReadConfig()
		if err != nil {
			log.Fatal("Failed to read config:", err)
		}
		if config.ChatGptApiKey == "" {
			fmt.Println("API key is not set")
		} else {
			fmt.Println("ChatGPT APIKey:", config.ChatGptApiKey)
		}
	},
}

func init() {
	rootCmd.AddCommand(setAPIKeyCmd)
	rootCmd.AddCommand(showAPIKeyCmd)
}
