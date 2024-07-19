package uglify

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/format"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var identifierMap = make(map[string]string)

func RandomizeCase(code string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())

	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Name != nil && !x.Name.IsExported() {
				randomizeName(x.Name)
			}
		case *ast.AssignStmt:
			for _, lhs := range x.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok && !ident.IsExported() {
					randomizeName(ident)
				}
			}
		case *ast.ValueSpec:
			for _, name := range x.Names {
				if !name.IsExported() {
					randomizeName(name)
				}
			}
		}
		return true
	})

	ast.Inspect(file, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && !ident.IsExported() {
			if newName, exists := identifierMap[ident.Name]; exists {
				ident.Name = newName
			}
		}
		return true
	})

	var buf strings.Builder
	err = format.Node(&buf, fset, file)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func randomizeName(ident *ast.Ident) {
	if _, exists := identifierMap[ident.Name]; !exists {
		identifierMap[ident.Name] = changeCaseRandomly(ident.Name)
	}
}

func changeCaseRandomly(str string) string {
	var modifiedStr strings.Builder
	for _, ch := range str {
		if rand.Float32() < 0.5 {
			modifiedStr.WriteRune(unicode.ToUpper(ch))
		} else {
			modifiedStr.WriteRune(unicode.ToLower(ch))
		}
	}
	return modifiedStr.String()
}
