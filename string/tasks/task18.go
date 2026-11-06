package tasks

import "strings"

func Join(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	var builder strings.Builder
	total := len(sep) * (len(parts) - 1)
	for _, part := range parts {
		total += len(part)
	}
	builder.Grow(total)
	for i, part := range parts {
		if i > 0 {
			builder.WriteString(sep)
		}
		builder.WriteString(part)
	}
	return builder.String()
}
