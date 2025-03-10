package submodules

import (
	"errors"
	"fmt"
	"math"
	"math/bits"

	"arit/modules"
	u "arit/modules/util"
)

var Prime = modules.Submodule{
	Name: "Prime",
	Keys: []string{"prime", "p", "Prime"},
	Help: "There is no help.",
}

func init() {
	check := &modules.Function{
		Name: "Prime number",
		Help: "Returns whether given number p is prime",
		N:    1,
		F: func(args []string) (any, error) {
			n, err := u.SingleInt64(args)
			if err != nil {
				return nil, err
			}
			return isprime(n)
		},
	}

	mersenne := &modules.Function{
		Name: "Mersenne",
		Help: "Returns whether given number p is a mersenne prime",
		N:    1,
		F: func(args []string) (any, error) {
			n, err := u.SingleInt64(args)
			if err != nil {
				return nil, err
			}
			return mersenne(n)
		},
	}

	factors := &modules.Function{
		Name: "Prime factors",
		Help: "Returns list of non-distinct aliquot parts of p",
		N:    1,
		F: func(args []string) (any, error) {
			n, err := u.SingleInt64(args)
			if err != nil {
				return nil, err
			}
			return factors2(n)
		},
	}

	funcs := map[string]*modules.Function{
		u.DEFAULT_KEY: check,
		"is":          check,
		"full":        check,
		"mersenne":    mersenne,
		"factors":     factors,
		"fac":         factors,
	}

	Prime.Funcs = funcs
}

// Returns whether given number p is prime
// should be AKS at some point
func isprime(p int64) (bool, error) {
	if p < 1 {
		return false, errors.New("negative numbers cannot be prime")
	}

	if p < 100 {
		return smallPrime[uint64(p)], nil
	}

	psq := int64(math.Ceil(math.Sqrt(float64(p))))

	var i int64 = 2

	for i < psq {
		if (p % i) == 0 {
			return isprime(i)
		}

		i++
	}

	return true, nil
}

// Returns whether given number p is a mersenne prime
func mersenne(p int64) (bool, error) {
	if p < 1 {
		return false, errors.New("negative numbers cannot be prime")
	}

	if p_ := (p + 1) & p; p_ != 0 {
		return false, nil
	}

	return isprime(p)
}

// Returns list of non-distinct aliquot parts of p
// Which is fancy talk for "prime factors"
func factors2(p int64) ([]uint32, error) {
	if p > 1<<32 {
		return []uint32{}, fmt.Errorf("ask someone else")
	}

	if p < 1 {
		return []uint32{}, fmt.Errorf("cannot factorize negative numbers")
	}

	if p < (1<<32)-1 {
		return _read(factorFile, factorTableFile, p)
	}

	return []uint32{}, fmt.Errorf("now yet impl for p larger than 2^16")

	lg2 := bits.LeadingZeros64(uint64(p))

	// i thought this was very clever
	factors := make([]uint32, lg2)
	idx := 0

	rem := p

	for rem&1 == 0 {
		factors[idx] = 2
		idx++
		rem >>= 1
	}

	for rem%2 == 0 {
		factors[idx] = 2
		idx++
		rem >>= 1
	}

	return factors, nil
}
