package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cocoide/commitify/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	configKey    = [...]string{"api-key", "language", "format"}
	configOption = [][]string{
		{},
		{"Japanese", "English"},
		{"Format 1", "Format 2"},
	}
)

type configModel struct {
	configKeyIndex    int
	configOptionIndex int
	configKeySelected bool
	err               error
	textInput         textinput.Model
}

func initConfigModel() configModel {
	ti := textinput.New()
	ti.Focus()

	return configModel{
		textInput: ti,
		err:       nil,
	}
}

func (cm configModel) Init() tea.Cmd {
	return textinput.Blink
}

func (cm configModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch cm.configKeySelected {
	// 設定項目を選択する
	case false:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyUp:
				if cm.configKeyIndex > 0 {
					cm.configKeyIndex--
				}
			case tea.KeyDown:
				if cm.configKeyIndex < len(configKey)-1 {
					cm.configKeyIndex++
				}
			case tea.KeyEnter:
				cm.configKeySelected = true
				return cm, nil
			case tea.KeyCtrlC, tea.KeyEsc:
				return cm, tea.Quit
			}
		}

	// 設定項目に値をセットする
	case true:
		switch len(configOption[cm.configKeyIndex]) {
		// 選択肢のない項目は入力を受け付ける
		case 0:
			var cmd tea.Cmd
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.Type {
				case tea.KeyEnter:
					saveConfig(cm)
					return cm, tea.Quit
				case tea.KeyCtrlC, tea.KeyEsc:
					return cm, tea.Quit
				}
			case error:
				cm.err = msg
				return cm, nil
			}

			cm.textInput, cmd = cm.textInput.Update(msg)
			return cm, cmd

		// 選択肢がある場合はセレクターで表示する
		default:
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.Type {
				case tea.KeyUp:
					if cm.configOptionIndex > 0 {
						cm.configOptionIndex--
					}
				case tea.KeyDown:
					if cm.configOptionIndex < len(configOption[cm.configKeyIndex])-1 {
						cm.configOptionIndex++
					}
				case tea.KeyEnter:
					saveConfig(cm)
					return cm, tea.Quit
				case tea.KeyCtrlC, tea.KeyEsc:
					return cm, tea.Quit
				}
			}
		}
	}

	return cm, nil
}

func (cm configModel) View() string {
	var b strings.Builder

	switch cm.configKeySelected {
	// 設定項目を選んでいない時
	case false:
		b.WriteString(color.WhiteString("設定項目を選んでください:\n"))
		b.WriteString(color.WhiteString("  ↑↓の矢印キーで項目を移動、Enterで選択\n"))

		for i, choice := range configKey {
			if i == cm.configKeyIndex {
				b.WriteString(fmt.Sprintf(color.HiCyanString("➡️  %s\n"), choice))
			} else {
				b.WriteString(fmt.Sprintf(color.CyanString("    %s\n"), choice))
			}
		}

	// 設定項目に値をセットする
	case true:
		// 選択肢のない項目はテキストエリアを表示
		switch len(configOption[cm.configKeyIndex]) {
		case 0:
			b.WriteString(color.WhiteString(fmt.Sprintf(
				"ここに%sを入力: %s\n",
				configKey[cm.configKeyIndex],
				cm.textInput.View(),
			)))
			b.WriteString(color.WhiteString("  Enterキーで確定"))

		default:
			b.WriteString(color.WhiteString("設定内容を選んでください:\n"))
			b.WriteString(color.WhiteString("  ↑↓の矢印キーで項目を移動、Enterで選択\n"))

			for i, option := range configOption[cm.configKeyIndex] {
				if i == cm.configOptionIndex {
					b.WriteString(fmt.Sprintf(color.HiCyanString("➡️  %s\n"), option))
				} else {
					b.WriteString(fmt.Sprintf(color.CyanString("    %s\n"), option))
				}
			}
		}
	}

	return b.String()
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "設定を変更します",
	Long:  `設定を変更します。設定項目はコマンドを実行すると表示されます。`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initConfigModel())
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func saveConfig(cm configModel) {
	currentConfig, err := util.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}

	switch cm.configKeyIndex {
	case 0:
		currentConfig.ChatGptApiKey = cm.textInput.Value()
	case 1:
		currentConfig.UseLanguage = configOption[cm.configKeyIndex][cm.configOptionIndex]
	case 2:
		currentConfig.CommitFormat = configOption[cm.configKeyIndex][cm.configOptionIndex]
	}

	err = util.WriteConfig(currentConfig)
	if err != nil {
		fmt.Println(err)
	}
}
