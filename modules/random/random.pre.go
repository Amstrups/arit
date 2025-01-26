package random

import (
	"errors"
	"strings"
)

func single(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New(" requires exactly 1 argument")
	}
	return args[0], nil
}

func trim(str string) string {
	return strings.TrimSuffix(strings.TrimPrefix(str, "\""), "\"")
}

