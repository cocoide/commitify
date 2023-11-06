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
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cocoide/commitify/internal/service"
)

var (
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

type suggestModel struct {
	choices    []string
	currentIdx int
	errorMsg   string
	isLoading  bool
	isEditing  bool
	spinner    spinner.Model
	textInput  textinput.Model
	scs        *service.SuggestCmdService
}

func (sm *suggestModel) Init() tea.Cmd {
	go func() {
		messages, err := sm.scs.GenerateCommitMessages()
		if err != nil {
			log.Fatal("コミットメッセージの生成に失敗: ", err)
			os.Exit(-1)
		}
		sm.choices = messages
		sm.isLoading = false
	}()

	return textinput.Blink
}

func (sm *suggestModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	sm.textInput, cmd = sm.textInput.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			sm.isEditing = true
			sm.textInput.Focus()
			sm.textInput.SetValue(sm.choices[sm.currentIdx])
			sm.textInput.CharLimit = 100
			sm.textInput.Width = 100
			return sm, cmd
		case tea.KeyUp:
			if sm.currentIdx > 0 {
				sm.currentIdx--
			}
		case tea.KeyDown:
			if sm.currentIdx < len(sm.choices)-1 {
				sm.currentIdx++
			}
		case tea.KeyEnter:
			if err := sm.scs.SubmitCommit(sm.choices[sm.currentIdx]); err != nil {
				sm.errorMsg = "コミットに失敗: " + err.Error()
				return sm, tea.Quit
			}
			return sm, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return sm, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		sm.spinner, cmd = sm.spinner.Update(msg)
		return sm, cmd
	}
	return sm, sm.spinner.Tick
}

func (sm *suggestModel) View() string {
	if sm.errorMsg != "" {
		return color.RedString(sm.errorMsg)
	}
	if sm.isLoading {
		s := fmt.Sprintf("\n %s %s\n\n", sm.spinner.View(), textStyle("コミットメッセージ生成中"))
		return s
	}
	var b strings.Builder
	if sm.errorMsg != "" {
		b.WriteString(color.RedString(sm.errorMsg) + "\n\n")
	}
	if sm.isEditing {
		return sm.textInput.View()
	}

	b.WriteString(color.WhiteString("🍕 Please select and enter to commit"))
	b.WriteString(color.WhiteString("\n  Use arrow ↑↓ to navigate and press Enter to select."))
	b.WriteString(color.WhiteString("\n  ( enter Tab key to edit message )\n\n"))

	for i, choice := range sm.choices {
		if i == sm.currentIdx {
			b.WriteString(fmt.Sprintf(color.HiCyanString("➡️  %s\n"), choice))
		} else {
			b.WriteString(fmt.Sprintf(color.CyanString("    %s\n"), choice))
		}
	}
	return b.String()
}

// モデルの初期化処理
func NewSuggestModel() *suggestModel {
	ti := textinput.New()
	ti.Focus()

	// suggestコマンドのサービスの取得
	scs, err := service.NewSuggestCmdService()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	return &suggestModel{
		choices:    []string{""},
		currentIdx: 0,
		errorMsg:   "",
		isLoading:  true,
		isEditing:  false,
		textInput:  ti,
		scs:        scs,
	}
}

// スピナーの初期化
func (sm *suggestModel) initSpinner() {
	sm.spinner = spinner.New()
	sm.spinner.Style = spinnerStyle
	sm.spinner.Spinner = spinner.Globe
}

var suggestCmd = &cobra.Command{
	Use:     "suggest",
	Short:   "Suggestion of commit message for staging repository",
	Aliases: []string{"s", "suggest"},
	Run: func(cmd *cobra.Command, args []string) {
		sm := NewSuggestModel()
		sm.initSpinner()
		p := tea.NewProgram(sm)
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
