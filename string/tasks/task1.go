package tasks

import "unicode/utf8"

func StringSize(s string) (bytes int, runes int) {
	return len(s), utf8.RuneCountInString(s)
}
