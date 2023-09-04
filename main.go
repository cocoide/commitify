package main

import (
	"fmt"
	"os"

	"github.com/cocoide/commitify/cmd"
)

func main() {
	// configファイルがあるかどうかを確認
	_, err := os.Stat("config.yaml")
	if os.IsNotExist(err) {
		if _, err := os.Create("config.yaml"); err != nil {
			fmt.Printf("error creating config file, %s", err.Error())
		}
	}

	cmd.Execute()
}
