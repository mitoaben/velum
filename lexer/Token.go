package lexer

import "velum/error"

type TokenType int

const (
	ID TokenType = iota

	STRING
	INT
	FLOAT
	HEX

	LET
	VAR
	CONST
	IF
	ELSE
	ELIF
	FOR
	RET
	TYPE
	FN
	AS
	RANGE
	IMPORT

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACK
	RBRACK

	ASSIGN
	STAR
	DIV
	ADD
	SUB

	COLON
	DOT
	COMMA
	NON

	LHS
	GTR
	SET
	FASSIGN
	NASSIGN
	LAND
	LOR
	LASSIGN
	MASSIGN
	DASSIGN
	SASSIGN
	AASSIGN

	EOF
	EOL
)

var TokenTypes = map[string]TokenType{
	"let":    LET,
	"var":    VAR,
	"const":  CONST,
	"if":     IF,
	"else":   ELSE,
	"elif":   ELIF,
	"for":    FOR,
	"type":   TYPE,
	"fn":     FN,
	"as":     AS,
	"range":  RANGE,
	"import": IMPORT,

	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
	"[": LBRACK,
	"]": RBRACK,

	"=": ASSIGN,
	"+": ADD,
	"-": SUB,
	"*": STAR,
	"/": DIV,
	"!": NON,

	":": COLON,
	".": DOT,
	",": COMMA,

	";": EOL,

	">>": LHS,
	"<<": GTR,
	"->": SET,
	":=": FASSIGN,
	"!=": NASSIGN,
	"&&": LAND,
	"||": LOR,
	"==": LASSIGN,
	"*=": MASSIGN,
	"/=": DASSIGN,
	"-=": SASSIGN,
	"+=": AASSIGN,
}

type Token struct {
	Type  TokenType
	Value string
	Pos   uint
	Line  uint
}

type Tokenizer struct {
	Er error.Error
}

func (tt TokenType) String() string {
	for name, t := range TokenTypes {
		if t == tt {
			return name
		}
	}
	return "UNKNOWN"
}
