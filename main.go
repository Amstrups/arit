package main

import (
  "arit/lexer_simple"
	"fmt"
)


func f(c chan int) {
  x := []int{1,2,3,4,5,6,7,8,9,0}
  for i, v := range x {
    if i % 3 == 0 {
      c <- v
    }
  }
  close(c)

}

func channelTesting() { 
  ch := make(chan int)
  go f(ch)
  for i := range ch {
      fmt.Println(i) 
      
  }
}

func main() { 
  channelTesting()
  return
  l := lexersimple.NewSLexer("+-*/\n111.111")
  for { 
    p,t,s := l.Lex()
    if t == lexersimple.EOF {
      return
    }
    fmt.Printf("Position: %d:%d, Token: %d, String: %s\n", p.Line, p.Column, t, s)
  }
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
