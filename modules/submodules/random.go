package submodules

import (
	"arit/modules"
	u "arit/modules/util"
	"fmt"
	"math/rand"
)

var Random = modules.Submodule{
	Name: "Random",
	Keys: []string{"Random", "rand", "random"},
	Help: "There is no help.",
}

func init() {
	number := &modules.Function{
		Name: "Number",
		Help: "Returns the most random number, sometimes",
		N:    0,
		F: func(args []string) (any, error) {
			if len(args) > 0 {
				return nil, fmt.Errorf("%s %s does not accept any args",
					"random", "num/number/default")
			}
			return number(), nil
		},
	}

	cap := &modules.Function{
		Name: "Capitilization",
		Help: "Returns the given string, with letter casing randomized",
		N:    1,
		F: func(args []string) (any, error) {
			str, err := u.Single(args)
			if err != nil {
				return nil, fmt.Errorf("%s::capitilization %s", Random.Name, err.Error())
			}
			return capitilization(str)
		},
	}

	gen := &modules.Function{
		Name: "Generate []int64",
		Help: "Given n, a and b, generates {[a,b)}^n",
		N:    3,
		F: func(args []string) (any, error) {
			n, a, b, err := u.TripleInt64(args)
			if err != nil {
				return []int64{}, fmt.Errorf("%s::generate %s", Random.Name, err.Error())
			}
			return generate64(n, a, b)
		},
	}

	closed := &modules.Function{
		Name: "Closed interval",
		Help: "Given a and b, generates x \\in [a,b]",
		N:    2,
		F: func(args []string) (any, error) {
			a, b, err := u.DoubleInt64(args)
			if err != nil {
				return []int64{}, fmt.Errorf("%s::closed_interval %s", Random.Name, err.Error())
			}
			return inOpen(a, b)
		},
	}

	funcs := map[string]*modules.Function{
		"number": number,
		"gen64":  gen,
		"cap":    cap,
		"closed": closed,
	}

	Random.Default = number
	Random.Funcs = funcs
}

// Returns the given string, with letter casing randomized
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

// Returns the most random number, sometimes
func number() int64 {
	x := rand.Uint64()
	var rem, xor uint64 = x, 0
	for range 64 {
		xor ^= rem & uint64(1)
		rem >>= 1
	}
	if xor == 1 {
		return int64((x << 1) >> 1)
	}

	return 17
}

// Given n, a and b, generates {[a,b)}^n
func generate64(n, a, b int64) ([]int64, error) {
	return _generate64(n, a, b)
}

// Given a and b, returns random number \in [a,b)
func inOpen(a, b int64) (int64, error) {
	if err := _verify(a, b); err != nil {
		return -1, err
	}

	return _verified(a, b), nil
}

// Given a and b, returns random number \in [a,b]
func inClosed(a, b int64) (int64, error) {
	b += 1
	if err := _verify(a, b); err != nil {
		return -1, err
	}

	return _verified(a, b), nil
}

// Return number in [1,a]
func d(a int64) (int64, error) {
	return inClosed(1, a)
}
