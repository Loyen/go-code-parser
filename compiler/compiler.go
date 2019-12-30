package compiler

import (
	"github.com/loyen/go-code-parser/lexer"
	"github.com/loyen/go-code-parser/parser"
)

// Compile by tokenizing source code and then processing the list of tokens
func Compile(source string) (string, error) {
	tokens, err := lexer.Tokenize(source)

	if err != nil {
		return "", err
	}

	code, err := parser.Parse(tokens)

	if err != nil {
		return "", err
	}

	return code, nil
}
