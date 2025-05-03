package parser

import (
	"velum/error"
	"velum/lexer"
)

type Parser struct {
	Root    *[]*Node
	Current *[]*Node
	Last    []*[]*Node
	Er      error.Error
}

type NodeClass int

const (
	UnkownNode NodeClass = iota
)

func NewParser(tokens *[]lexer.Token, er error.Error) Parser {
	var parser Parser
	var newNodeRoot []*Node

	parser.Er = er
	parser.Current = &newNodeRoot
	parser.Root = &newNodeRoot
	parser.Last = append(parser.Last, &newNodeRoot)

	return parser
}

func (p *Parser) Parse(tokens *[]lexer.Token) *[]*Node {

	cop := []lexer.Token{}
	var line uint = 1

	save := func() {
		p.mast(&cop)
		cop = cop[:0]
	}

	for _, token := range *tokens {
		if token.Line != line {
			if len(cop) != 0 {
				save()
			}
			line = token.Line
		} else if token.Type == lexer.EOF || token.Type == lexer.EOL {
			if len(cop) != 0 {
				save()
			}
		}
		if token.Type != lexer.EOF && token.Type != lexer.EOL {
			cop = append(cop, token)
		}
	}

	return p.Root
}

var OperatorPrecedence = map[string]int{
	"=":  1,
	"||": 2,
	"&&": 3,

	"==": 4, "!=": 4,
	"<<": 4, "<=": 4, ">>": 4, ">=": 4,

	"+": 5, "-": 5,
	"*": 6, "/": 6, "%": 6,

	"^": 7,
}

type Node struct {
	Class NodeClass
	Value any
}

type ExprNode struct {
	Left  *Node
	Right *Node
	Op    lexer.TokenType
}

type FnNode struct {
	Name       string
	Args       *[]*Node
	Body       *[]*Node
	ReturnType *Node
}

type CallNode struct {
	Name string
	Args *[]*Node
}

type UCondNode struct {
	Init lexer.TokenType
	Cond *Node
	Body *[]*Node
}
