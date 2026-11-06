package tasks

import "strings"

func RemoveRune(s string, target rune) string {
	var builder strings.Builder
	for _, r := range s {
		if r != target {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
