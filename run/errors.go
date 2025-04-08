package run

import (
	"strings"
)

/*
	Builtin joinError concatting errors without carriage return
*/

type joinedError struct {
	errors []error
}

func (joined joinedError) Error() (err string) {
	for _, e := range joined.errors {
		err += e.Error()
	}

	return strings.TrimRight("\n", err)
}

func JoinErrors(base error, right error) error {
	if base == nil {
		return right
	}

	if right == nil {
		return base
	}

	if j_base, ok := base.(joinedError); ok {
		j_base.errors = append(j_base.errors, right)
		return j_base
	}

	joined := joinedError{
		errors: []error{base},
	}

	if j_base, ok := right.(joinedError); ok {
		joined.errors = append(joined.errors, j_base.errors...)
		return joined
	}

	joined.errors = append(joined.errors, right)
	return joined
}
