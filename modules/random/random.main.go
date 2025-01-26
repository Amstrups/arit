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
PRE { this.trim }
*/
func capitilization(str string) (string, error) {
	B := []byte(str)
	B_ := make([]byte, len(B))

	for i, b := range B {
		B_[i] = b
		if b < 'A' || b > 'z' || (b > 'Z' && b < 'a') {
			continue
		}

		if rand.Intn(2) > 0 {
			// flips bit indicating whether upper- or lowercase
			B_[i] ^= 32
		}
	}
	return string(B_), nil
}

/*
FUNCTION: { num, number }
HELP: { Returns the most random number }
*/
func number() string {
	return "17"
}

/*
FUNCTION: { gen, gen64, generate64 }
HELP: { Given n, a and b, generates {[a,b)}^n }
PRE { TripleInt64 }
*/
func generate64(n, a, b int64) ([]int64, error) {
	return _generate64(n, a, b)
}

/*
FUNCTION: { inopen, open }
HELP: { Given a and b, returns random number \inOpen [a,b) }
PRE { DoubleInt64 }
*/
func inOpen(a, b int64) (int64, error) {
	if err := _verify(a, b); err != nil {
		return -1, err
	}

	return _verified(a, b), nil
}

/*
FUNCTION: { in, inclosed, closed }
HELP: { Given a and b, returns random number \inOpen [a,b] }
PRE { DoubleInt64 }
*/
func inClosed(a, b int64) (int64, error) {
	b += 1
	if err := _verify(a, b); err != nil {
		return -1, err
	}

	return _verified(a, b), nil
}

/*
FUNCTION: { d6 }
HELP: { Return number in [1,6] }
*/
func d6() (int64, error) {
	return inClosed(1, 6)
}

/*
FUNCTION: { d3 }
HELP: { Return number in [1,6] }
*/
func d3() (int64, error) {
	return inClosed(1, 3)
}
