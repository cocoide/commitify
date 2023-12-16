/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cocoide/commitify/internal/entity"
	"github.com/cocoide/commitify/internal/gateway"
	"github.com/cocoide/commitify/internal/usecase"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"strings"
)

type pushModel struct {
	step            PushCmdStep
	selectBaseIndex int
	baseList        []string
	pr              *entity.PullRequest
	pcu             *usecase.PushCmdUsecase
	loadMsg         string
	errMsg          string
	spinner         spinner.Model
	prInput         *prInput
}

type prInput struct {
	titleInput textinput.Model
	bodyInput  textarea.Model
}

type PushCmdStep int

const (
	SelectBaseBranch PushCmdStep = iota
	EditPRTitle
	EditPRBody
	SubmitPR
)

var _ tea.Model = &pushModel{}

func (m *pushModel) Init() tea.Cmd {
	return nil
}

type generatePRMsg struct {
	pr  *entity.PullRequest
	err error
}

func (m *pushModel) generatePRCmd() tea.Cmd {
	return func() tea.Msg {
		selectBranch := m.baseList[m.selectBaseIndex]
		pr, err := m.pcu.GeneratePullRequest(selectBranch)
		return generatePRMsg{pr, err}
	}
}

type submitPRMsg struct {
	err error
}

func (m *pushModel) submitPRCmd() tea.Cmd {
	return func() tea.Msg {
		err := m.pcu.SubmitPullRequest(m.pr)
		return submitPRMsg{err}
	}
}

func (m *pushModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyCtrlE:
			m.step--
		}
		switch m.step {
		case SelectBaseBranch:
			switch msg.Type {
			case tea.KeyUp:
				if m.selectBaseIndex > 0 {
					m.selectBaseIndex--
				}
			case tea.KeyDown:
				if m.selectBaseIndex < len(m.baseList)-1 {
					m.selectBaseIndex++
				}
			case tea.KeyEnter:
				m.loadMsg = "PRを生成中..."
				return m, m.generatePRCmd()
			}
		case EditPRTitle:
			cmd = m.updateTitleInput(msg)
			switch msg.Type {
			case tea.KeyEnter:
				m.pr.Title = m.prInput.titleInput.Value()
				m.step = EditPRBody
				m.focusInPRBody()
				return m, cmd
			}
		case EditPRBody:
			cmd = m.updateBodyInput(msg)
			switch msg.Type {
			case tea.KeyEnter:
				m.pr.Body = m.prInput.bodyInput.Value()
				m.step = SubmitPR
				m.loadMsg = "PRを提出中..."
				return m, m.submitPRCmd()
			}
		}
	case generatePRMsg:
		if msg.err != nil {
			m.errMsg = fmt.Sprintf("🚨PR生成中にエラーが発生: %v", msg.err)
		} else {
			m.pr = msg.pr
			m.step = EditPRTitle
			m.focusInPRTitle()
		}
		m.finishLoading()
	case submitPRMsg:
		if msg.err != nil {
			// リファクタ: github tokenがexpireした時は、loginコマンドを自動実行
			// → tokenがexpireしたときのエラーを調べないといけない...
			m.errMsg = fmt.Sprintf("🚨PR提出中にエラーが発生: %v", msg.err)
		}
		m.finishLoading()
	}
	return m, nil
}

func (m *pushModel) View() string {
	if m.errMsg != "" {
		return m.errMsg
	}
	if m.loadMsg != "" {
		return fmt.Sprintf("\n %s %s\n\n", m.spinner.View(), textStyle(m.loadMsg))
	}
	switch m.step {
	case SelectBaseBranch:
		return m.buildSelectBaseBranchText()
	case EditPRTitle:
		return color.HiWhiteString("<Titleの編集: Enterで確定>\n\n") + fmt.Sprintf("Title:\n%s", m.prInput.titleInput.View())
	case EditPRBody:
		return color.HiWhiteString("<Bodyの編集: EnterでPush&PR提出, Ctrl+Eで戻る>\n\n") + fmt.Sprintf("Body:\n%s", m.prInput.bodyInput.View())
	case SubmitPR:
		return fmt.Sprintf("**🎉PRの作成に成功**")
	}
	return ""
}

func NewPushModel() *pushModel {
	ti := textinput.New()
	ti.Focus()
	var errMsg string
	ctx := context.Background()
	github := gateway.NewGithubGateway()
	nlp := gateway.NewOpenAIGateway(ctx)
	pcu := usecase.NewPushCmdUsecase(github, nlp)
	baseList, err := pcu.GetRemoteBaseBranchCandidates()
	if err != nil {
		errMsg = "🚨Baseブランチを取得中にエラーが発生"
	}
	var maxBaseElements = 5
	var selectedBaseList []string
	if len(baseList) < maxBaseElements {
		selectedBaseList = baseList[:]
	} else {
		selectedBaseList = baseList[:maxBaseElements]
	}
	return &pushModel{
		step:     SelectBaseBranch,
		baseList: selectedBaseList,
		pcu:      pcu,
		errMsg:   errMsg,
		loadMsg:  "",
		prInput: &prInput{
			titleInput: textinput.New(),
			bodyInput:  textarea.New(),
		},
	}
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push and create pull request",
	Run: func(cmd *cobra.Command, args []string) {
		m := NewPushModel()
		m.initSpinner()
		p := tea.NewProgram(m)
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func (m *pushModel) buildSelectBaseBranchText() string {
	var b strings.Builder
	b.WriteString(color.HiWhiteString("📢Baseブランチ(Merge先)を選んで下さい\n"))
	b.WriteString(color.WhiteString("click ↑↓ to navigate and press Enter to select.\n"))

	for i, base := range m.baseList {
		if i == m.selectBaseIndex {
			b.WriteString(fmt.Sprintf(color.HiCyanString("➡️  %s\n"), base))
		} else {
			b.WriteString(fmt.Sprintf(color.CyanString("    %s\n"), base))
		}
	}
	return b.String()
}

func (m *pushModel) finishLoading() {
	m.loadMsg = ""
}

func (m *pushModel) initSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.Globe
}

func (m *pushModel) focusInPRTitle() {
	input := m.prInput.titleInput
	input.Focus()
	input.SetValue(m.pr.Title)
	input.CharLimit = 100
	input.Width = 100
	m.prInput.titleInput = input
}

func (m *pushModel) focusInPRBody() {
	input := m.prInput.bodyInput
	input.Focus()
	input.SetValue(m.pr.Body)
	input.CharLimit = 5000
	input.SetWidth(200)
	m.prInput.bodyInput = input
}

func (m *pushModel) updateTitleInput(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.prInput.titleInput, cmd = m.prInput.titleInput.Update(msg)
	return cmd
}

func (m *pushModel) updateBodyInput(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.prInput.bodyInput, cmd = m.prInput.bodyInput.Update(msg)
	return cmd
}
