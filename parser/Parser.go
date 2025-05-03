package parser

import (
	"fmt"
	"velum/lexer"
)

func (p Parser) mast(tokens *[]lexer.Token) {
	for _, token := range *tokens {
		fmt.Printf("%+v\n", token)
	}
}
