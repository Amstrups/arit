package parser

import l "arit/parsing/lexer"

type Parser struct {
	lexer *l.Lexer
}

func New(lexer *l.Lexer) *Parser {
	return &Parser{lexer}
}

func (p *Parser) Parse() string {
	return "fuck"
}
