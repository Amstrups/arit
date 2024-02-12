package main

import (
	lexer "arit/lexer_simple"
	"arit/modes/prime"
	parser "arit/parser_simple"
	"flag"
	"fmt"
)


func main() { 
  var primeFlag = flag.Int("p", -1, "check if input is primenumber")
  if primeFlag == nil {
    fmt.Println("nothing to do")
    return
  }
  flag.Parse()
  prime.IsPrime(*primeFlag) 


  return 
  l := lexer.NewSLexer("+-*/\n111.111")
  p := parser.NewParser(l)
  stmt := p.Parse()
  fmt.Println(stmt)
}

/* Scanner example

    var s scanner.Scanner
    input := "x := 5 + 3\nif x > 5 {\n    fmt.Println(\"x is greater than 5\")\n}" 
    s.Init(strings.NewReader(input))
    for {
        tok := s.Scan()
        if tok == scanner.EOF {
            break
        }
        fmt.Printf("Token: %s, Value: %s\n", s.TokenText(), string(tok))
    }
*/

/* Args example 

  fmt.Println(l)
  return 
  argsWithProg := os.Args 

  flag.Parse()
  fmt.Println(argsWithProg[1])
*/
