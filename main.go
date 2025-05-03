package main

import (
	"os"
	"velum/compiler"
)

func main() {

	var c compiler.Compiler

	mode := 0

	for _, arg := range os.Args[1:] {
		switch mode {
		case 0:
			switch arg {
			case "-o":
				mode = 1
			case "-f":
				mode = 2
			default:
				c.InputFile = arg
			}
		case 1:
			c.OutputFile = arg
			mode = 0
		case 2:
			c.Plarform = arg
			mode = 0
		}
	}

	c.Compile()
}
