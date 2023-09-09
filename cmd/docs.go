/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
		color.Cyan(string(b))
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
