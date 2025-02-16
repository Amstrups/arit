package submodules

import (
	"errors"
	"fmt"

	u "arit/modules/util"
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

func (p *Prime) Parse(Args []string) (any, error) {
	if len(Args) == 0 {
		return nil,
			fmt.Errorf("Prime module does not have non-arg default function")
	}
	f := Args[0]
	args := Args[1:]

	switch f {
	case "full":
		n, err := u.SingleInt64(args)
		if err != nil {
			return nil, err
		}
		return p.isprime(n)
	case "is":
		n, err := u.SingleInt64(args)
		if err != nil {
			return nil, err
		}
		return p.mersenne(n)
	default:
		n, err := u.SingleInt64(Args)
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

	if p <= 3 && p > 1 {
		return true, nil
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
func (*Prime) mersenne(p int64) (bool, error) {
	if p_ := (p + 1) & p; p_ != 0 {
		return false, nil
	}

	if p < 1 {
		return false, errors.New("negative numbers cannot be prime")
	}

	if p <= 3 && p > 1 {
		return true, nil
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

// Returns list of non-distinct aliquot parts of p
// Which is fancy talk for "prime factors"
func (*Prime) factors(p int64) (bool, error) {
	if p < 1 {
		return false, errors.New("negative numbers cannot be prime")
	}

	if p <= 3 && p > 1 {
		return true, nil
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
