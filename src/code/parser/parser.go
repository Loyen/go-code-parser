package parser

import (
	"bytes"
	"code/lexer"
	"strconv"
)

var buffer bytes.Buffer
var debugBuffer bytes.Buffer
var tokens []lexer.Token

// Parse a list of tokens and return the processed output
func Parse(tokenList []lexer.Token) string {
	tokens = tokenList

	for len(tokens) > 0 {
		statement()
	}

	buffer.WriteString("\nDEBUGGER\n" + debugBuffer.String())

	return buffer.String()
}

func statement() {
	token := shiftTokens()

	switch token.Type {
		case "print":
			statementPrint()
		default:
			expression()
	}
}

func statementPrint() {
	buffer.WriteString(expression())
}

func expression() string {
	token := shiftTokens()

	switch token.Type {
		case "string":
			return token.Value
		default:
			return ""
	}
}


func shiftTokens() lexer.Token {
	var token lexer.Token

	token = tokens[0]
	tokens = tokens[1:]

	debugToken(token)

	return token
}

func debugToken(token lexer.Token) {
	debugBuffer.WriteString("L" + strconv.Itoa(token.Line) + ":")
	debugBuffer.WriteString("T" + token.Type + ":")
	debugBuffer.WriteString(token.Value)
	debugBuffer.WriteString("\n")
}
