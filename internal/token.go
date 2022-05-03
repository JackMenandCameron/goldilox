package internal

import (
	"fmt"
)

type Token struct {
	Ttype  TokenType
	Lexeme string
	Value string
	Line int
}

func MakeToken(ttype TokenType, line int) Token {
	return Token{Ttype: ttype, Line: line}
}

func (t Token) String() string {
	return fmt.Sprintf("[%s: %s]", TTToString[t.Ttype], t.Lexeme)
}
