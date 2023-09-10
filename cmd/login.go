/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type loginModel struct {
	DeviceCodeResponse deviceCodeResponse
}

type deviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationUri string `json:"verification_uri"`
	ExpiresIn       string `json:"interval"`
}

func initLoginModel() loginModel {
	// POSTリクエスト
	values := url.Values{}
	values.Set("client_id", "aadc21dd58f5aa8db2f0")
	values.Add("scope", "read:user")

	endPoint := "https://github.com/login/device/code"

	req, err := http.NewRequest("POST", endPoint, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	res := sendRequest(req)

	// レスポンスの検証
	var responseData deviceCodeResponse
	err = json.Unmarshal(res, &responseData)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("URLにアクセスして、次の認証コードを入力してください\n")
	fmt.Printf("URL: %v\n", responseData.VerificationUri)
	fmt.Printf("認証コード: %v\n", responseData.UserCode)

	for

	// lm = loginModel{}
	// lm.device_code =

	return loginModel{}
}

func (lm loginModel) Init() tea.Cmd {
	return nil
}

func (lm loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return lm, tea.Quit
		}
	}

	return lm, nil
}

func (lm loginModel) View() string {
	var b strings.Builder

	b.WriteString(color.BlackString("test login."))

	return b.String()
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "ログイン処理のテスト実装",
	Long:  `ログイン処理のテスト実装です。`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initLoginModel())
		p.Run()
	},
}

func sendRequest(req *http.Request) []byte {
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error")
	}

	return byteArray
}
