package tasks

import (
	"strings"
	"unicode"
)

func NormalizeSpaces(s string) string {
	var builder strings.Builder
	inSpace := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			if builder.Len() > 0 {
				inSpace = true
			}
			continue
		}
		if inSpace {
			builder.WriteByte(' ')
			inSpace = false
		}
		builder.WriteRune(r)
	}
	return builder.String()
}
