package parser

import (
	"bytes"
	"errors"
	"github.com/loyen/go-code-parser/lexer"
	"strconv"
)

var buffer bytes.Buffer
var debugBuffer bytes.Buffer
var tokens []lexer.Token

// Parse a list of tokens and return the processed output
func Parse(tokenList []lexer.Token) (string, error) {
	tokens = tokenList

	var err error

	for len(tokens) > 0 {
		err = statement()

		if err != nil {
			return "", err
		}
	}

	return buffer.String(), nil
}

func statement() error {
	token, err := shiftTokens()

	if err != nil {
		return err
	}

	switch token.Type {
		case "print":
			err := statementPrint()

			if err != nil {
				return err
			}
		case "command_end":
			return nil
		default:
			return errors.New("Unknown token \"" + token.Type + "\"")
	}

	return nil
}

func statementPrint() error {
	expression, err := expression()

	if err != nil {
		return err
	}

	err = expectToken("command_end")
	shiftTokens()

	if (err != nil) {
		return err
	}

	buffer.WriteString(expression)
	return nil
}

func expression() (string, error) {
	token, err := shiftTokens()

	if err != nil {
		return "", err
	}

	switch token.Type {
		case "string":
			return token.Value, nil
		case "number":
			return token.Value, nil
		default:
			return "", nil
	}
}

func expectToken(tokenType string) error {
	if (len(tokens) < 1) {
		return errors.New("Unexpected end of file")
	}

	var token lexer.Token
	token = tokens[0]

	if (token.Type == tokenType) {
		return nil
	}

	return errors.New("Unexpected token: " + token.Type + ". Expected \"" + tokenType + "\"")
}


func shiftTokens() (lexer.Token, error) {
	var token lexer.Token

	if len(tokens) < 1 {
		return token, errors.New("Unexpected end of file")
	}

	token = tokens[0]
	tokens = tokens[1:]

	debugToken(token)

	return token, nil
}

func debugToken(token lexer.Token) {
	debugBuffer.WriteString("L" + strconv.Itoa(token.Line) + ":")
	debugBuffer.WriteString("T" + token.Type + ":")
	debugBuffer.WriteString(token.Value)
	debugBuffer.WriteString("\n")
}
