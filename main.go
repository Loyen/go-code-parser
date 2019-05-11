package main

import (
	"fmt"
	"github.com/loyen/code_parser/compiler"
)

func main() {
	source := "print \"Hello world!\"\n"

	code := compiler.Compile(source)

	fmt.Printf(code)
}
