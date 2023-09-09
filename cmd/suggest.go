package cmd

import (
	"context"
	"fmt"
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
	isLoading  bool
	animationIdx int
	messages   []string
}

type generateMessages struct {
	messages []string
	errorMsg string
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		og := gateway.NewOpenAIGateway(ctx)
		ms := service.NewMessageService(og)
		stagingCode := util.ExecGetStagingCode()
		messages, err := ms.GenerateCommitMessage(stagingCode)
		if err != nil {
			return generateMessages{errorMsg: "メッセージの生成に失敗: " + err.Error()}
		}
		return generateMessages{messages: messages}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case generateMessages:
		if msg.errorMsg != "" {
			m.errorMsg = msg.errorMsg
			m.isLoading = false
			return m, nil
		}
		m.choices = msg.messages
		m.isLoading = false
		return m, nil
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
	case tea.Cmd:
		if m.isLoading {
			m.animationIdx = (m.animationIdx + 1) % 3
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.errorMsg != "" {
		red := color.New(color.FgRed).SprintFunc()
		return fmt.Sprintf(red(m.errorMsg))
	}
	if m.isLoading {
		AnimationEarth := []string{"🌎","🌍","🌏"}
		AnimationPoint := []string{".","..","..."}
		return fmt.Sprintf("%s Generating commit messages %s", AnimationEarth[m.animationIdx], AnimationPoint[m.animationIdx])
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
		p := tea.NewProgram(m)
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
