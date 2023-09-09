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
	Short: "Document of committify",
	Run: func(cmd *cobra.Command, args []string) {
		b, _ := static.Logo.ReadFile("logo.txt")
		cyan := color.New(color.FgCyan).SprintFunc()
		logo := cyan(string(b))
		fmt.Println(logo)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
