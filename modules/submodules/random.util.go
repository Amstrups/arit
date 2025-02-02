package submodules

import (
	"errors"
	"math/rand"
)

func _abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func _verify(a, b int64) error {
	if a >= b {
		return errors.New("must satisfy: a < b")
	}
	return nil
}

func _verified(a, b int64) int64 {
	return rand.Int63n(b-_abs(a)) + a
}

func _generate64(n, a, b int64) (X []int64, err error) {
	if err = _verify(a, b); err != nil {
		return X, err
	}

	if n == 0 {
		err = errors.New("must satisfy: n > 0")
		return
	}

	X = make([]int64, n)
	for i := range X {
		X[i] = _verified(a, b)
	}

	return
}
