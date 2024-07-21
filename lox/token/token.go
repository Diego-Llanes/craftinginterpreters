package token

import "fmt"

type Token struct {
	TokenType int 
	lexeme string
	literal interface{}
	line int
}

func NewToken(
		tokenType int,
		lexeme string,
		literal interface{},
		line int,
	) *Token{
	out := Token{}
	out.TokenType = tokenType
	out.lexeme = lexeme
	out.literal = literal
	out.line = line
	return &out
}

func (t *Token) ToString () string {
	return string(t.TokenType) + " " + t.lexeme + " " + fmt.Sprintf("%v", t.literal)
}

// func (t *Token) Fn (source string) {}
