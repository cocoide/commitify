package util

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecGetStagingCode() string {
	code, err := exec.Command("git", "diff", "--staged").Output()
	if err != nil {
		fmt.Printf("Gitでエラーが発生")
		log.Fatal(err.Error())
	}
	return string(code)
}

func ExecCommitMessage(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
