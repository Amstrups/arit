package main

import (
	"github.com/amstrups/nao"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	println("hello world")

	spew.Dump(nao.Run("2+2;0b011x8"))
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
