package main

import (
	"fmt"
	"code/compiler"
)

func main() {
	source := "print \"Hello world!\"\n"

	code := compiler.Compile(source)

	fmt.Printf(code)
}
