package tasks

import "sort"

func TopKWords(text string, k int) []string {
	if k <= 0 {
		return []string{}
	}
	frequencies := WordFrequency(text)
	words := make([]string, 0, len(frequencies))
	for word := range frequencies {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		fi, fj := frequencies[words[i]], frequencies[words[j]]
		if fi != fj {
			return fi > fj
		}
		return words[i] < words[j]
	})
	if k > len(words) {
		k = len(words)
	}
	return words[:k]
}
