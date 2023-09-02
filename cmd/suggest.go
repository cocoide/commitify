package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/cocoide/commitify/internal/gateway"
	"github.com/cocoide/commitify/internal/service"
	"github.com/cocoide/commitify/util"
	"github.com/spf13/cobra"
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggestion of commit message for staging repository",
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
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
