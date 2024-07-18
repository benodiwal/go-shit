package utils

import (
	"log"
	"os"
)

func ReadFile(filename string) (content []byte) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	return content
}
