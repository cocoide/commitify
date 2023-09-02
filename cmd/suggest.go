package cmd

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/cocoide/commitify/internal/gateway"
	"github.com/cocoide/commitify/util"
	"github.com/spf13/cobra"
)

const (
	CommitMessagePrompt = "Generate commit message for [%s]"
	FormatNotice        = ", format your commit as:\n- feat: [feature description]\n- bugfix: [bugfix description]"
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggestion of commit message for staging repository",
	Run: func(cmd *cobra.Command, args [] string) {
		util.LoadEnv()
		ctx := context.Background()
		og := gateway.NewOpenAIGateway(ctx)
		result, err := exec.Command("git", "diff", "--staged").Output()
		if err != nil {
			log.Fatal(err.Error())
		}
		// 設定に応じてPromptは動的に変化させる
		prompt := fmt.Sprintf(CommitMessagePrompt, string(result))
		answer, err := og.GetAnswerFromPrompt(prompt, 0.01)
		// 生成中のUIを非同期で表示
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(answer)
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
