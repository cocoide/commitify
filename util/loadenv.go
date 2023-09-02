package util

import (
	"log"

	"github.com/joho/godotenv"
)

// Viperを導入したらgodotenvのコードは削除する
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("failed to load .env file: %v" + err.Error())
	} else {
		// log.Print("success to load .env file")
	}
}
