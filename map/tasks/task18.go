package tasks

import "unicode"

func AreAnagrams(a, b string) bool {
	counts := make(map[rune]int)
	for _, r := range a {
		if !unicode.IsSpace(r) {
			counts[unicode.ToLower(r)]++
		}
	}
	for _, r := range b {
		if unicode.IsSpace(r) {
			continue
		}
		key := unicode.ToLower(r)
		counts[key]--
		if counts[key] < 0 {
			return false
		}
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}
