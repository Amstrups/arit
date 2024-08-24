package util

import "errors"

type T any
type TS []T

var ErrStackEmpty = errors.New("stack is empty")

func (s TS) Push(i int) TS {
	return append(s, i)
}
func (s TS) Pop() (TS, T, error) {
	l := len(s)
	if l == 0 {
		return s, -1, ErrStackEmpty
	}
	return s[:l-1], s[l-1], nil
}

func (s TS) IsEmpty() bool {
	return len(s) == 0
}
