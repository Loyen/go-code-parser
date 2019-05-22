package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/loyen/code_parser/compiler"
)

func main() {
	var filepath string

	if len(os.Args) < 1 {
		fmt.Println("No file inputted")
		return
	}

	filepath = os.Args[1]

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println("File could not be read")
		return
	}

	source := string(content)

	code, err := compiler.Compile(source)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(code)
}
