package lexersimple

type Token int

const (
  EOF = iota
	ILLEGAL
	IDENT
	INT
  FLOAT
	SEMI

	ADD
	SUB
	MUL
	DIV
  POW

	ASSIGN

  LPAREN
  RPAREN
)

