package main

import (
	"arit/parsing"
	lexer "arit/parsing/lexer"
	parser "arit/parsing/parser"
	"fmt"
)

func main() {
	l := lexer.New("2+2")
	for {
		tok := l.Lex()
		if tok.T == parsing.EOF {
			return

		}
		fmt.Println(tok.T)

	}
	p := parser.New(l)
	stmt := p.Parse()
	fmt.Println(stmt)
}

/*
  var primeFlag = flag.Int("p", -1, "check if input is primenumber")
  if primeFlag == nil {
    fmt.Println("nothing to do")
    return
  }
  flag.Parse()
  prime.IsPrime(*primeFlag)
*/
