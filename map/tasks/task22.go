package tasks

func GroupByLength(words []string) map[int][]string {
	result := make(map[int][]string)

	for _, word := range words {

		length := 0
		for range word {
			length++
		}

		result[length] = append(result[length], word)
	}

	return result
}
