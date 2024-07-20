package lox

type Token struct {
	TokenType int 
	lexeme string
	literal interface{}
	line int
}

func (f *Token) Fn (source string) {}


