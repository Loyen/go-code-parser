package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/loyen/go-code-parser/compiler"
)

func main() {
	var filepath string

	if len(os.Args) < 2 {
		fmt.Println("No file inputted")
		return
	}

	filepath = os.Args[1]

	output, err := RunFile(filepath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(output)
}

func RunFile(filepath string) (string, error) {
	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	source := string(content)

	output, err := Run(source)

	if err != nil {
		return "", err
	}

	return output, nil
}

func Run(source string) (string, error) {
	output, err := compiler.Compile(source)

	if err != nil {
		return "", err
	}

	return output, nil
}
