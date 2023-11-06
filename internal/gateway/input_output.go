package gateway

import (
	"github.com/cocoide/commitify/internal/service"
	"os/exec"
)

type inputOutputGateway struct {
}

func NewInputOutputGateway() service.GithubService {
	return &inputOutputGateway{}
}

func (g *inputOutputGateway) GetStaginCodeDiff() (string, error) {
	// Gitが入ってるかどうかのチェックも入れる
	// 入っていないなら専用のエラーメッセージを生成
	diff, err := exec.Command("git", "diff", "--staged").Output()
	return string(diff), err
}
