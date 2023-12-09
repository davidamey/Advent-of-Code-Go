package util

import "strings"

func Divide(s, sep string) (string, string) {
	parts := strings.SplitN(s, sep, 2)
	return parts[0], parts[1]
}
