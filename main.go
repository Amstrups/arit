package main

import (
	"arit/parser"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() { 

  argsWithProg := os.Args 



  //input := "x := 5 + 3\nif x > 5 {\n fmt.Println(\"x is greater than 5\")\n}"
  //input := "..."
  tokens := parser.Lexer([]byte(argsWithProg[1])) 
  tree := parser.Parser(tokens)
  result, err := tree[0].Eval()
  if (err != nil) {
    log.Fatal(err)
  }
  println("eval: ", result)
  

  //name := flag.String("name", "world", "The name to greet.")
  flag.Parse()
  fmt.Println(argsWithProg[1])
}

/*
package main
import (
    "fmt"
    "text/scanner"
    "strings"
)
func main() {
    var s scanner.Scanner
    input := "x := 5 + 3\nif x > 5 {\n    fmt.Println(\"x is greater than 5\")\n}" s.Init(strings.NewReader(input))
    for {
        tok := s.Scan()
        if tok == scanner.EOF {
            break
        }
        fmt.Printf("Token: %s, Value: %s\n", s.TokenText(), string(tok))
    }
}
*/
