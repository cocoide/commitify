/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type loginModel struct {
}

func initLoginModel() loginModel {
	return loginModel{}
}

func (lm loginModel) Init() tea.Cmd {
	return nil
}

func (lm loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return lm, tea.Quit
		}
	}

	return lm, nil
}

func (lm loginModel) View() string {
	var b strings.Builder

	b.WriteString(color.BlackString("test login."))

	return b.String()
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "ログイン処理のテスト実装",
	Long:  `ログイン処理のテスト実装です。`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initLoginModel())
		p.Run()
	},
}
