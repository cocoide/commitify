package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commitify",
	Short: "CLI for thinking commit message",
	Long: `By "commitify config" command, you can change commit message format or language, 
	( To know details about format or language, enter commitify docs )`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
