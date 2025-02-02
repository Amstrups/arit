package submodules

import (
	"errors"
)

// go:generate go run ./cmd/generatos/module.go
// MODULE: p, prime

/*
FUNCTION: is
HELP: Returns whether given number p is prime
PRE: modules.SingleInt64
*/
func isprime(p int64) (bool, error) {
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
