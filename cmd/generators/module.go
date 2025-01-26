//go:build generate

package main

import (
	"os"
)

func main() {
	f, _ := os.Create("hello-world.md")
	f.Close()
}
