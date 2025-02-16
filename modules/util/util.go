package preprocess

import (
	"arit/cli/parser"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Single(args []parser.Argument) (string, error) {
	if len(args) != 1 {
		return "", errors.New(" requires exactly 1 argument")
	}

	return args[0].V, nil
}

func Trim(str string) string {
	return strings.TrimSuffix(strings.TrimPrefix(str, "\""), "\"")
}

func SingleInt64(args []parser.Argument) (int64, error) {
	if len(args) != 1 {
		return -1, errors.New(" requires exactly 1 argument")
	}

	if args[0].T != parser.NUMBER {
		return -1, fmt.Errorf("expected number, found %v", args[0].T)
	}

	return strconv.ParseInt(args[0].V, 10, 64)
}

func DoubleInt64(args []parser.Argument) (a, b int64, err error) {
	if len(args) != 3 {
		err = errors.New(" requires exactly 3 argument")
		return
	}

	for _, a := range args {
		if a.T != parser.NUMBER {
			return -1, -1, fmt.Errorf("expected number, found %v", a.T)
		}
	}

	a, err = strconv.ParseInt(args[0].V, 10, 64)
	if err != nil {
		return
	}

	b, err = strconv.ParseInt(args[1].V, 10, 64)
	return
}

func TripleInt64(args []parser.Argument) (a, b, c int64, err error) {
	if len(args) != 3 {
		err = errors.New(" requires exactly 3 argument")
		return
	}

	for _, a := range args {
		if a.T != parser.NUMBER {
			return -1, -1, -1, fmt.Errorf("expected number, found %v", a.T)
		}
	}

	a, err = strconv.ParseInt(args[0].V, 10, 64)
	if err != nil {
		return
	}

	b, err = strconv.ParseInt(args[1].V, 10, 64)
	if err != nil {
		return
	}

	c, err = strconv.ParseInt(args[2].V, 10, 64)
	if err != nil {
		return
	}
	return
}

func KInt64s(args []parser.Argument, k int) (parsed []int64, err error) {
	parsed = make([]int64, k)

	for i := 0; i < k; i++ {
		if args[i].T != parser.NUMBER {
			err = fmt.Errorf("expected number, found %v", args[i].T)
			return
		}

		parsed[i], err = strconv.ParseInt(args[i].V, 10, 64)
	}
	return
}
