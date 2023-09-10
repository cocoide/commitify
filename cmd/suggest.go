package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/gateway"
	"github.com/cocoide/commitify/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type model struct {
	choices    []string
	currentIdx int
	errorMsg   string
	isLoading  bool
}

func (m *model) Init() tea.Cmd {
	conf, err := entity.ReadConfig()
	if err != nil {
		log.Fatal("設定情報の取得に失敗: ", err)
	}

	var gi gateway.GatewayInterface
	switch conf.AISource {
	case int(entity.WrapServer):
		gi = gateway.NewGrpcServeGateway()
	default:
		gi = gateway.NewGrpcServeGateway()
	}

	messages, err := gi.FetchCommitMessages()
	if err != nil {
		log.Fatal("コミットメッセージの生成に失敗: ", err)
		os.Exit(-1)
	}
	m.choices = messages
	m.isLoading = false

	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.errorMsg = "コミットに失敗: " + err.Error()
				return m, tea.Quit
			}
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *model) View() string {
	if m.errorMsg != "" {
		red := color.New(color.FgRed).SprintFunc()
		return fmt.Sprintf(red(m.errorMsg))
	}
	if m.isLoading {
		return "🌎 Generating commit messages ..."
	}
	var b strings.Builder
	if m.errorMsg != "" {
		red := color.New(color.FgRed).SprintFunc()
		b.WriteString(red(m.errorMsg) + "\n\n")
	}
	white := color.New(color.FgWhite).SprintFunc()
	b.WriteString(white("🍕Please select an option:"))
	b.WriteString(white("\n  Use arrow ↑↓ to navigate and press Enter to select.\n\n"))

	for i, choice := range m.choices {
		cyan := color.New(color.FgCyan).SprintFunc()
		hiCyan := color.New(color.FgHiCyan).SprintFunc()
		if i == m.currentIdx {
			b.WriteString(fmt.Sprintf(hiCyan("➡️  %s\n"), choice))
		} else {
			b.WriteString(fmt.Sprintf(cyan("    %s\n"), choice))
		}
	}
	return b.String()
}

var suggestCmd = &cobra.Command{
	Use:     "suggest",
	Short:   "Suggestion of commit message for staging repository",
	Aliases: []string{"s", "suggest"},
	Run: func(cmd *cobra.Command, args []string) {
		m := model{isLoading: true}
		p := tea.NewProgram(&m)
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
