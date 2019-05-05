package compiler

import (
	"code/lexer"
	"code/parser"
)

// Compile by tokenizing source code and then processing the list of tokens
func Compile(source string) string {
	tokens := lexer.Tokenize(source)
	code := parser.Parse(tokens)

	return code
}
