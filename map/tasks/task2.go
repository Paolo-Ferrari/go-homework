package tasks

func LastIndex(words []string) map[string]int {
	result := make(map[string]int)
	for i, word := range words {
		result[word] = i
	}
	return result
}
