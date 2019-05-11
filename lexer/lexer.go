package lexer

import (
	"bytes"
	"unicode"
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
func Tokenize(source string) []Token {

	mode = modeDefault

	line := 1

	runes := []rune(source)
	for _, character := range runes {
		if string(character) == "\n" {
			line++
		}

		switch mode {
			case modeDefault:
				processModeCode(character, line)
			case modeString:
				processModeString(character, line)
		}
	}

	buffer.Reset()

	return tokens
}

func processModeCode(character rune, line int) {
	if !unicode.IsSpace(character) {
		buffer.WriteString(string(character))
	}

	if buffer.String() == "print" {
		token := Token{ Line: line, Type: "print", Value: "" }
		tokens = append(tokens, token)
		buffer.Reset()

		return
	} else if string(character) == "\"" {
		buffer.Reset()

		mode = modeString
		return
	}
}

func processModeString(character rune, line int) {
	if string(character) == "\"" {
		token := Token{ Line: line, Type: "string", Value: buffer.String() }
		tokens = append(tokens, token)

		buffer.Reset()

		mode = modeDefault
		return
	}

	buffer.WriteString(string(character))
}
