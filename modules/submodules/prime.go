package submodules

import (
	"errors"
	"fmt"
	"math/bits"

	"arit/cli/parser"
	u "arit/modules/util"
)

const (
	factorFile      = "./bin/file"
	factorTableFile = "./bin/table"
)

type Prime struct{}

func (*Prime) Name() string {
	return "Prime"
}
func (*Prime) Keys() []string {
	return []string{"prime", "p"}
}
func (*Prime) Description() string {
	return "Module for primeness in arit"
}

func (p *Prime) Parse(cmd parser.Command) (any, error) {
	switch cmd.Func {
	case "is", "full":
		n, err := u.SingleInt64(cmd.Args)
		if err != nil {
			return nil, err
		}
		return p.isprime(n)

	case "mersenne":
		n, err := u.SingleInt64(cmd.Args)
		if err != nil {
			return nil, err
		}
		return p.isprime(n)
	case "factors", "fac":
		n, err := u.SingleInt64(cmd.Args)
		if err != nil {
			return nil, err
		}
		return p.factors2(n)
	default:
		n, err := u.SingleInt64(cmd.Args)
		if err != nil {
			return nil, err
		}
		return p.isprime(n)
	}
}

// Returns whether given number p is prime
func (*Prime) isprime(p int64) (bool, error) {
	if p < 1 {
		return false, errors.New("negative numbers cannot be prime")
	}

	if p < 100 {
		return smallPrime[uint64(p)], nil
	}

	if p <= 1 || p%2 == 0 || p%3 == 0 {
		return false, nil
	}

	for i := int64(5); i*i <= p; i += 6 {
		if p%i == 0 || p%(i+2) == 0 {
			return false, nil
		}
	}
	return true, nil
}

// Returns whether given number p is a mersenne prime
func (mod *Prime) mersenne(p int64) (bool, error) {
	if p < 1 {
		return false, errors.New("negative numbers cannot be prime")
	}

	if p_ := (p + 1) & p; p_ != 0 {
		return false, nil
	}

	return mod.isprime(p)
}

// Returns list of non-distinct aliquot parts of p
// Which is fancy talk for "prime factors"
func (*Prime) factors2(p int64) ([]uint32, error) {
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
