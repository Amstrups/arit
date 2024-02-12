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

func Name(t Token) string {
  switch t {
  case EOF:
    return "EOF"
  case ILLEGAL:
    return "ILLEGAL"
  case IDENT:
    return "identifier"
  case INT:
    return "int"
  case FLOAT:
    return "float"
  case ADD:
    return "+"
  case SUB:
    return "-"
  case MUL:
    return "*"
  case DIV:
    return "\\"
  case POW:
    return "^"
  case LPAREN:
    return "("
  case RPAREN:
    return ")"
  default:
    return "Not a (used) token"
  }
}
