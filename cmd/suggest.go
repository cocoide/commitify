package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cocoide/commitify/internal/gateway"
	"github.com/cocoide/commitify/internal/service"
	"github.com/cocoide/commitify/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type model struct {
	choices    []string
	currentIdx int
	errorMsg   string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.currentIdx > 0 {
				m.currentIdx--
			}
		case tea.KeyDown:
			if m.currentIdx < len(m.choices)-1 {
				m.currentIdx++
			}
		case tea.KeyEnter:
			if err := util.ExecCommitMessage(m.choices[m.currentIdx]); err != nil {
				m.errorMsg = "コミットエラーが発生"
				return m, tea.Quit
			}
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	if m.errorMsg != "" {
		red := color.New(color.FgRed).SprintFunc()
		b.WriteString(red(m.errorMsg) + "\n\n")
	}
	b.WriteString("Please select an option:\n\n")

	for i, choice := range m.choices {
		blue := color.New(color.FgCyan).SprintFunc()
		if i == m.currentIdx {
			b.WriteString(fmt.Sprintf(blue("-> %s\n"), choice))
		} else {
			b.WriteString(fmt.Sprintf("   %s\n", choice))
		}
	}

	b.WriteString("\nUse the arrow keys to navigate and press Enter to select.")
	return b.String()
}

var suggestCmd = &cobra.Command{
	Use:     "suggest",
	Short:   "Suggestion of commit message for staging repository",
	Aliases: []string{"s", "suggest"},
	Run: func(cmd *cobra.Command, args []string) {
		util.LoadEnv()
		ctx := context.Background()
		og := gateway.NewOpenAIGateway(ctx)
		ms := service.NewMessageService(og)
		msgCh, err := ms.AsyncGenerateCommitMessage()
		if err != nil {
			log.Fatal(err.Error())
		}
		suggestMsg := <-msgCh
		fmt.Println(suggestMsg)
		choices := []string{suggestMsg, "!", "!#!"}
		m := model{choices: choices}
		p := tea.NewProgram(m)
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
