package random

import (
	"math/rand"
)

/*
MODULE: r, rand, random
*/

/*
FUNCTION { cap, capitilization }
HELP { Returns the given string, with letter casing randomized }
PRE { trim }
*/
func capitilization(str string) string {
	B := []byte(str)
	B_ := make([]byte, len(B))

	for i, b := range B {
		B_[i] = b
		if b < 'A' || b > 'z' || (b > 'Z' && b < 'a') {
			continue
		}

		if rand.Intn(2) > 0 {
			B_[i] ^= 32 // flips bit indicating upper- or lowercase
		}
	}
	return string(B_)
}

/*
FUNCTION: { num, number }
HELP: { Returns the most random number }
*/
func number() string {
	return "17"
}
