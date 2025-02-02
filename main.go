package main

import (
	"arit/cli"
	"os"
)

func main() {
	err := cli.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
}
