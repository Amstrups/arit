package submodules

import (
	"arit/modules"
	u "arit/modules/util"
	"fmt"
)

var Economics = modules.Submodule{
	Name: "Economics",
	Keys: []string{"econ", "economics", "Random"},
	Help: "There is no help.",
}

func init() {
	terms := &modules.Function{
		Name: "Terms",
		Help: "Returns amount of terms expected before principal is payed off",
		N:    3,
		F: func(args []string) (any, error) {
			principal, payment, err := u.DoubleInt64(args[:2])
			if err != nil {
				return 0, err
			}

			interest, err := u.SingleFloat64(args[2:])
			if err != nil {
				return 0, err
			}

			return terms(principal, payment, interest)
		},
	}

	funcs := map[string]*modules.Function{
		u.DEFAULT_KEY: terms,
		"terms":       terms,
	}

	Economics.Funcs = funcs
}

// Return expected amount of terms until loan is payed off,
// assuming interest are applying before payment
func terms(principal, payment int64, interest float64) (int64, error) {
	if interest == 0 || interest > 300 {
		return 0, fmt.Errorf("not supporting interest rates of 300%%, given %f.4", interest)
	}

	if principal < 0 {
		return 0, fmt.Errorf("principal less than 0 are not supported ")
	}

	if payment < 0 {
		return 0, fmt.Errorf("payments less than 0 are not supported ")
	}

	if principal/payment > 400 {
		return 0, fmt.Errorf("not gonna happen")
	}

	rem := float64(principal)
	var x int64

	for rem > 0 {
		rem *= interest
		rem -= float64(payment)
		x++
	}

	return x, nil

}
