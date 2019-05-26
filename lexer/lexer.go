package lexer

import (
	"bytes"
	"unicode"
	"errors"
)

type Token struct {
	Line int
	Type string
	Value string
}


const modeDefault = 1
const modeString = 2

var mode int

var buffer bytes.Buffer

var tokens []Token

// Tokenize a string of code
func Tokenize(source string) ([]Token, error) {

	mode = modeDefault

	line := 1

	runes := []rune(source)
	for _, character := range runes {
		if string(character) == "\n" {
			line++
		}

		switch mode {
			case modeDefault:
				err := processModeCode(character, line)

				if err != nil {
					return nil, err
				}
			case modeString:
				err := processModeString(character, line)

				if err != nil {
					return nil, err
				}
			default:
				return nil, errors.New("Unknown lexer mode")
		}
	}

	if (len(tokens) == 0 || tokens[len(tokens)-1].Type != "command_end") {
		token := Token{ Line: line, Type: "command_end", Value: "" }
		tokens = append(tokens, token)
	}
	buffer.Reset()

	return tokens, nil
}

func processModeCode(character rune, line int) error {
	if !unicode.IsSpace(character) {
		buffer.WriteString(string(character))
	}

	if string(character) == "\n" {
		token := Token{ Line: line, Type: "command_end", Value: "" }
		tokens = append(tokens, token)

		return nil
	} else if buffer.String() == "print" {
		token := Token{ Line: line, Type: "print", Value: "" }
		tokens = append(tokens, token)
		buffer.Reset()

		return nil
	} else if string(character) == "\"" {
		buffer.Reset()

		mode = modeString

		return nil
	}

	return nil
}

func processModeString(character rune, line int) error {
	if string(character) == "\"" {
		token := Token{ Line: line, Type: "string", Value: buffer.String() }
		tokens = append(tokens, token)

		buffer.Reset()

		mode = modeDefault

		return nil
	}

	buffer.WriteString(string(character))

	return nil
}
