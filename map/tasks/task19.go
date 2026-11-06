package tasks

import "unicode"

func WordFrequency(text string) map[string]int {
	result := make(map[string]int)
	var word []rune
	flush := func() {
		if len(word) == 0 {
			return
		}
		result[string(word)]++
		word = word[:0]
	}
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			word = append(word, unicode.ToLower(r))
		} else {
			flush()
		}
	}
	flush()
	return result
}
