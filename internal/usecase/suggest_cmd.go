package usecase

import (
	"fmt"
	"github.com/cocoide/commitify/internal/service"
	"log"
	"os/exec"

	"github.com/cocoide/commitify/internal/entity"
)

type SuggestCmdUsecase struct {
	message service.CommitMessageService
	github  service.GithubService
}

func NewSuggestCmdUsecase(message service.CommitMessageService, github service.GithubService) *SuggestCmdUsecase {
	conf, err := entity.ReadConfig()
	if err != nil {
		log.Fatalf("設定の読み込みに失敗")
	}
	if conf.WithGptRequestLocation() == entity.Local {
		log.Fatalf("現在、非対応の機能です。")
	}
	return &SuggestCmdUsecase{message: message, github: github}
}

func (u *SuggestCmdUsecase) GenerateCommitMessages() ([]string, error) {
	stagingCodeDiff, err := u.github.GetStaginCodeDiff()
	// stagingCodeを取捨選択する処理をここに入れる
	if err != nil {
		return nil, err
	}
	conf, err := entity.ReadConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to open config file: %v", err)
	}
	return u.message.GenerateCommitMessageList(stagingCodeDiff, conf)
}

func (u *SuggestCmdUsecase) SubmitCommit(commitMessage string) error {
	cmd := exec.Command("git", "commit", "-m", commitMessage)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
