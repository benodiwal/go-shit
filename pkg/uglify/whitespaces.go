package uglify

import (
	"go/scanner"
	"go/token"
	"math/rand"
	"strings"
	"time"
)

func AddRandomWhiteSpace(code string) string {
	var modifiedSrc strings.Builder
	rand.Seed(time.Now().UnixNano())
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(code))
	s.Init(file, []byte(code), nil, scanner.ScanComments)

	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		if isPunctuator(tok) {
			randomWhitespaceBefore := getRandomWhitespace()
			modifiedSrc.WriteString(randomWhitespaceBefore)
		}

		if lit != "" {
			modifiedSrc.WriteString(lit)
		} else {
			modifiedSrc.WriteString(" ")
		}

		if isPunctuator(tok) {
			randomWhitespaceAfter := getRandomWhitespace()
			modifiedSrc.WriteString(randomWhitespaceAfter)
		}
	}

	return modifiedSrc.String()
}

func isPunctuator(tok token.Token) bool {
	switch tok {
	case token.LPAREN, token.RPAREN, token.LBRACE, token.RBRACE,
		token.LBRACK, token.RBRACK, token.COMMA, token.PERIOD,
		token.SEMICOLON, token.COLON, token.ASSIGN, token.DEFINE,
		token.ARROW, token.ELLIPSIS:
		return true
	}
	return false
}

func getRandomWhitespace() string {
	whitespaceOptions := []string{" ", "  ", "\t", "\t\t", "\t "}
	return whitespaceOptions[rand.Intn(len(whitespaceOptions))]
}
