package tasks

import "unicode/utf8"

func FirstRune(s string) (rune, bool) {
	if s == "" {
		return 0, false
	}
	r, _ := utf8.DecodeRuneInString(s)
	return r, true
}
