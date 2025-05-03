package lexer

import (
	"fmt"
	"unicode"
)

var Line uint = 1
var Start uint = 0

func (t Tokenizer) Tokenize(src *[]byte) *[]Token {
	input := *src
	var tokens []Token
	length := len(input)
	pos := 0

	for pos < length {
		Start++
		ch := input[pos]

		// Пропуск пробелов и управляющих символов
		if isWhitespace(ch) {
			pos++
			continue
		}

		// Однострочный комментарий //
		if ch == '/' && pos+1 < length && input[pos+1] == '/' {
			pos += 2
			for pos < length && input[pos] != '\n' {
				pos++
			}
			continue
		}

		// Многострочный комментарий /* ... */
		if ch == '/' && pos+1 < length && input[pos+1] == '*' {
			pos += 2
			for pos < length-1 {
				if input[pos] == '*' && input[pos+1] == '/' {
					pos += 2
					break
				}
				pos++
			}
			continue
		}

		start := pos

		// Идентификаторы и ключевые слова
		if isAlpha(ch) || ch == '_' {
			for pos < length && (isAlnum(input[pos]) || input[pos] == '_') {
				pos++
			}
			value := string(input[start:pos])
			typ, ok := TokenTypes[value]
			if !ok {
				typ = ID
			}
			tokens = append(tokens, Token{Type: typ, Value: value, Pos: uint(Start), Line: Line})
			continue
		}

		// Числа: INT, FLOAT, HEX
		if isDigit(ch) {
			isFloat := false
			if ch == '0' && pos+1 < length && (input[pos+1] == 'x' || input[pos+1] == 'X') {
				// Hex
				pos += 2
				for pos < length && isHexDigit(input[pos]) {
					pos++
				}
				tokens = append(tokens, Token{Type: HEX, Value: string(input[start:pos]), Pos: uint(Start), Line: Line})
			} else {
				// Int or Float
				for pos < length && isDigit(input[pos]) {
					pos++
				}
				if pos < length && input[pos] == '.' {
					isFloat = true
					pos++
					for pos < length && isDigit(input[pos]) {
						pos++
					}
				}
				tokType := INT
				if isFloat {
					tokType = FLOAT
				}
				tokens = append(tokens, Token{Type: tokType, Value: string(input[start:pos]), Pos: uint(Start), Line: Line})
			}
			continue
		}

		// Строки
		if ch == '"' {
			pos++
			for pos < length && input[pos] != '"' {
				if input[pos] == '\\' && pos+1 < length {
					pos += 2
				} else {
					pos++
				}
			}
			if pos < length && input[pos] == '"' {
				pos++
			}
			tokens = append(tokens, Token{Type: STRING, Value: string(input[start:pos]), Pos: uint(Start), Line: Line})
			continue
		}

		// Операторы и скобки
		found := false
		for l := 3; l >= 1; l-- {
			if pos+l <= length {
				slice := input[pos : pos+l]
				if typ, ok := TokenTypes[string(slice)]; ok {
					tokens = append(tokens, Token{Type: typ, Value: string(slice), Pos: uint(Start), Line: Line})
					pos += l
					found = true
					break
				}
			}
		}
		if found {
			continue
		}

		t.Er.ReportError(fmt.Sprintf("Unknown symbol %c", ch), Line, uint(start))
		pos++
	}

	return &tokens
}

func isWhitespace(ch byte) bool {
	if ch == '\n' {
		Line++
		Start = 0
	}
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isAlpha(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isAlnum(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch))
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isHexDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9') ||
		('a' <= ch && ch <= 'f') ||
		('A' <= ch && ch <= 'F')
}
