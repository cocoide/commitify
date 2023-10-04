package main

import (
	"fmt"
	"os"

	"github.com/cocoide/commitify/cmd"
)

func main() {
	// configファイルがあるかどうかを確認
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error of find user home dir, %v", err)
		return
	}

	_, err = os.Stat(homePath + "/.commitify/config.yaml")
	if os.IsNotExist(err) {
		if err := os.MkdirAll(homePath+"/.commitify", 0755); err != nil {
			fmt.Printf("error of make directory, %v", err)
		}
		if _, err := os.Create(homePath + "/.commitify/config.yaml"); err != nil {
			fmt.Printf("error creating config file, %s", err.Error())
		}
	}

	cmd.Execute()
}
