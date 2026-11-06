package tasks

import "strings"

func CollapseRuns(s string) string {
	var builder strings.Builder
	var previous rune
	first := true
	for _, r := range s {
		if first || r != previous {
			builder.WriteRune(r)
			previous = r
			first = false
		}
	}
	return builder.String()
}
