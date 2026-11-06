package tasks

import "unicode"

func TrimUnicodeSpace(s string) string {
	runes := []rune(s)
	start, end := 0, len(runes)
	for start < end && unicode.IsSpace(runes[start]) {
		start++
	}
	for start < end && unicode.IsSpace(runes[end-1]) {
		end--
	}
	return string(runes[start:end])
}
