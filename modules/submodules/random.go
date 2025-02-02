package submodules

import (
	u "arit/modules/util"
	"fmt"
	"math/rand"
)

type Random struct{}

func (*Random) Name() string { return "Random" }
func (*Random) Description() string {
	return "Module for randomess in arit"
}

func (r *Random) Parse(Args []string) (any, error) {
	if len(Args) == 0 {
		return r.number(), nil
	}
	f := Args[0]
	args := Args[1:]

	switch f {
	case "num", "number":
		if len(args) > 0 {
			return nil, fmt.Errorf("subcommand %s %s does not accept any args", "random", f)

		}
		return r.number(), nil
	case "gen", "gen64", "generate64":
		n, a, b, err := u.TripleInt64(args)
		if err != nil {
			return nil, err
		}
		return r.generate64(n, a, b)
	case "cap", "capitilization":
		str, err := u.Single(args)
		if err != nil {
			return nil, err
		}
		return r.capitilization(str)
	default:
		return nil, nil
	}
}

func (*Random) Help() string {
	return "There is no help."
}

// Returns the given string, with letter casing randomized
func (*Random) capitilization(str string) (string, error) {
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
func (*Random) number() int64 {
	x := rand.Uint64()
	var rem, xor uint64 = x, 0
	for i := 0; i < 64; i++ {
		xor ^= rem & uint64(1)
		rem >>= 1
	}

	return 17
}

// Given n, a and b, generates {[a,b)}^n
func (*Random) generate64(n, a, b int64) ([]int64, error) {
	return _generate64(n, a, b)
}

// Given a and b, returns random number \in [a,b)
func (*Random) inOpen(a, b int64) (int64, error) {
	if err := _verify(a, b); err != nil {
		return -1, err
	}

	return _verified(a, b), nil
}

// Given a and b, returns random number \in [a,b]
func (*Random) inClosed(a, b int64) (int64, error) {
	b += 1
	if err := _verify(a, b); err != nil {
		return -1, err
	}

	return _verified(a, b), nil
}

// Return number in [1,a]
func (r *Random) d(a int64) (int64, error) {
	return r.inClosed(1, a)
}
