package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/gateway"
	"github.com/cocoide/commitify/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)

type model struct {
	choices    []string
	currentIdx int
	errorMsg   string
	isLoading  bool
	isEditing  bool
	spinner    spinner.Model
	textInput  textinput.Model
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

	go func() {
		messages, err := gi.FetchCommitMessages()
		if err != nil {
			log.Fatal("コミットメッセージの生成に失敗: ", err)
			os.Exit(-1)
		}
		m.choices = messages
		m.isLoading = false
	}()

	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			m.isEditing = true
			m.textInput.Focus()
			m.textInput.SetValue(m.choices[m.currentIdx])
			m.textInput.CharLimit = 100
			m.textInput.Width = 100
			return m, cmd
		case tea.KeyUp:
			if m.currentIdx > 0 {
				m.currentIdx--
			}
		case tea.KeyLeft:
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
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, m.spinner.Tick
}

func (m *model) resetSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.Globe
}

func (m *model) View() string {
	if m.errorMsg != "" {
		return color.RedString(m.errorMsg)
	}
	if m.isLoading {
		s := fmt.Sprintf("\n %s %s\n\n", m.spinner.View(), textStyle("コミットメッセージ生成中"))
		return s
	}
	var b strings.Builder
	if m.errorMsg != "" {
		b.WriteString(color.RedString(m.errorMsg) + "\n\n")
	}
	if m.isEditing {
		return m.textInput.View()
	}

	b.WriteString(color.WhiteString("🍕 Please select and enter to commit"))
	b.WriteString(color.WhiteString("\n  Use arrow ↑↓ to navigate and press Enter to select."))
	b.WriteString(color.WhiteString("\n  ( enter Tab key to edit message )\n\n"))

	for i, choice := range m.choices {
		if i == m.currentIdx {
			b.WriteString(fmt.Sprintf(color.HiCyanString("➡️  %s\n"), choice))
		} else {
			b.WriteString(fmt.Sprintf(color.CyanString("    %s\n"), choice))
		}
	}
	return b.String()
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()

	return model{
		choices:    []string{""},
		currentIdx: 0,
		errorMsg:   "",
		isLoading:  true,
		isEditing:  false,
		textInput:  ti,
	}
}

var suggestCmd = &cobra.Command{
	Use:     "suggest",
	Short:   "Suggestion of commit message for staging repository",
	Aliases: []string{"s", "suggest"},
	Run: func(cmd *cobra.Command, args []string) {
		m := initialModel()
		m.resetSpinner()
		p := tea.NewProgram(&m)
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
