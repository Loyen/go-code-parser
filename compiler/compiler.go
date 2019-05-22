package compiler

import (
	"github.com/loyen/code_parser/lexer"
	"github.com/loyen/code_parser/parser"
)

// Compile by tokenizing source code and then processing the list of tokens
func Compile(source string) (string, error) {
	tokens := lexer.Tokenize(source)
	code, err := parser.Parse(tokens)

	if err != nil {
		return "", err
	}

	return code, nil
}
