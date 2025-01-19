package random

import (
	"math/rand"
)

func offset(b byte, chain int) (byte, int) {
	if b < 'A' || b > 'z' || (b > 'Z' && b < 'a') {
		return b, 0
	}

	n := 20

	x := rand.Intn(n+chain) - n

	if x < int(n/2) {
		return b & 223, 0
	}

	return b | 255, 0
}

func Capitilization2(str string) string {

	B := []byte(str)
	B_ := make([]byte, len(B))
	chain := 0

	for i, b := range B {
		B_[i], chain = offset(b, chain)
	}
	return string(B_)
}
