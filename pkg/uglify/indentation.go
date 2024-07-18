package uglify

import (
	"math/rand"
	"strings"
	"time"
)

func RandomIndentation(code string) string {
	rand.Seed(time.Now().UnixNano())
	
	lines := strings.Split(code, "\n")
	var uglifiedLines []string
	for _, line := range lines {
		indentation := strings.Repeat("  ", rand.Intn(10))
		uglifiedLines = append(uglifiedLines, indentation+line)
	}
	uglifiedContent := strings.Join(uglifiedLines, "\n")
	return uglifiedContent
}
