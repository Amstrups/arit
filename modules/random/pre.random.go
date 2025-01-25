package random

import "strings"

func trim(str string) string {
	return strings.TrimSuffix(strings.TrimPrefix(str, "\""), "\"")
}
