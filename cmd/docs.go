/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cocoide/commitify/static"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Document of commitify",
	Run: func(cmd *cobra.Command, args []string) {
		b, _ := static.Logo.ReadFile("logo.txt")
		fmt.Print(color.CyanString(string(b)) + "\n\n ・Languageは日本語と英語が選択できます\n\n ・CodeFormatはPrefix (例: feat: A)とEmoji (例: 🐛 Bugix), Normal (例: Feat A)が選べます")
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
