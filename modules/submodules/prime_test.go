package submodules

import (
	"fmt"
	"math"
	"os/exec"
	"testing"
)

var list []int64

func init() {
	r := Random{}
	l, err := r.generate64(600, 10000, 30000)
	if err != nil {
		panic(err)
	}

	list = l
}

func BenchmarkSqrtBc(b *testing.B) {
	outs := make([]int64, len(list))
	for i, e := range list {
		str := fmt.Sprintf("\"sqrt(%d)\"", e)

		out, err := exec.Command("bc", "-e", str).Output()

		if err != nil {
			panic(err)
		}

		var val int64 = 0
		start := min(8, len(out))
		for j := start; j > 0; j-- {
			val |= int64(out[start-j]) << j * 8
		}
		outs[i] = val
	}
}

func BenchmarkSqrtBuilt(b *testing.B) {
	outs := make([]int64, len(list))
	for i, e := range list {
		outs[i] = int64(math.Sqrt(float64(e)))
	}
}
