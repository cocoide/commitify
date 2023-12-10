/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/cocoide/commitify/internal/gateway"
	"github.com/cocoide/commitify/internal/usecase"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"time"
)

const (
	DeviceActivateURL = "https://github.com/login/device"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login by github",
	Long:  `by login you can use auto pull request feature`,
	Run: func(cmd *cobra.Command, args []string) {
		httpClient := gateway.NewHttpClient()
		u := usecase.NewLoginCmdUsecase(httpClient)
		res, err := u.BeginGithubSSO()
		if err != nil {
			fmt.Printf("ログイン中にエラーが発生: %v", err)
		}
		go func() {
			req := &usecase.ScheduleVerifyAuthRequest{
				DeviceCode: res.DeviceCode, Interval: res.Interval, ExpiresIn: res.ExpiresIn}
			if err := u.ScheduleVerifyAuth(req); err != nil {
				fmt.Printf("error occured: %v", err)
			}
		}()
		fmt.Printf("以下のページで認証コード『%s』を入力して下さい。\n", res.UserCode)
		fmt.Printf(color.HiCyanString("➡️  %s\n"), DeviceActivateURL)
		time.Sleep(time.Second * time.Duration(res.ExpiresIn))
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
