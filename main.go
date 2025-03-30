package main

import (
	"arit/run"
	"os"
)

func main() {
	err := run.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
}
