package main

import (
	"arit/cli"
	"os"
)

func main() {
	//fmt.Println(os.Args[1:])
	err := cli.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
}
