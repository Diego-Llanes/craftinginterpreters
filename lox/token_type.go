package lox

const (
	// Single Charachter tokens
	Left_Paren = iota
	Right_Paren
	Left_Brace
	Right_Brace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// One to two characther tokens
	Bang
	Bang_Equal
	Equal
	Equal_Equal
	Greater
	Greater_Equal
	Less
	Less_Equal

	// Literals
	Identifier
	String
	Number

	// Keywords
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	EOF
)

