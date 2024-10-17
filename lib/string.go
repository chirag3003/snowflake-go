package lib

import "strings"

func BuildString(len int, str ...string) string {
	var builder strings.Builder
	builder.Grow(len)
	for _, s := range str {
		builder.WriteString(s)
	}
	return builder.String()
}
