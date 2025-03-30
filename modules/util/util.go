package util

import (
	"errors"
	"strconv"
	"strings"
)

const DEFAULT_KEY = "DEFAULT_KEY"

func Single(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New(" requires exactly 1 argument")
	}
	return args[0], nil
}

func Trim(str string) string {
	return strings.TrimSuffix(strings.TrimPrefix(str, "\""), "\"")
}

func SingleInt64(args []string) (int64, error) {
	if len(args) != 1 {
		return -1, errors.New(" requires exactly 1 argument")
	}
	return strconv.ParseInt(args[0], 10, 64)

}

func DoubleInt64(args []string) (a, b int64, err error) {
	if len(args) != 2 {
		err = errors.New(" requires exactly 2 argument")
		return
	}

	a, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return
	}

	b, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return
	}
	return
}

func TripleInt64(args []string) (a, b, c int64, err error) {
	if len(args) != 3 {
		err = errors.New(" requires exactly 3 argument")
		return
	}

	a, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return
	}

	b, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return
	}

	c, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return
	}
	return
}

func KInt64s(args []string, k int) (parsed []int64, err error) {
	parsed = make([]int64, k)

	for i := range k {
		parsed[i], err = strconv.ParseInt(args[i], 10, 64)
	}
	return
}

func SingleFloat64(args []string) (float64, error) {
	if len(args) != 1 {
		return -1, errors.New(" requires exactly 1 argument")
	}
	return strconv.ParseFloat(args[0], 32)
}
