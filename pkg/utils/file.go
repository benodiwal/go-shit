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

func WriteFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatalf("failed to write file: %s", err)
	}
}
