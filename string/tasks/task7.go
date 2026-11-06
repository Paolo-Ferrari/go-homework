package tasks

import "strings"

func ReplaceRune(s string, old, new rune) string {
	var builder strings.Builder
	for _, r := range s {
		if r == old {
			r = new
		}
		builder.WriteRune(r)
	}
	return builder.String()
}
