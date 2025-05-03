package compiler

import (
	"fmt"
	"os"
	"velum/error"
	"velum/lexer"
	"velum/parser"
)

type Compiler struct {
	OutputFile string
	InputFile  string
	Plarform   string
}

func (c Compiler) Compile() {
	compileFile(c.InputFile)
}

func compileFile(file string) {
	src, err := os.ReadFile(file)
	if err != nil {
		println("Read file \"" + file + "\" error")
		os.Exit(1)
	}

	var tokenizer lexer.Tokenizer
	var er error.Error
	er.Modul = file
	tokenizer.Er = er

	tokens := tokenizer.Tokenize(&src)
	*tokens = append(*tokens, lexer.Token{Type: lexer.EOF})

	// Create and run the parser
	p := parser.NewParser(tokens, er)
	program := p.Parse(tokens)

	fmt.Println(program)

}
