package tasks

func MergeCounts(a, b map[string]int) map[string]int {
	result := make(map[string]int, len(a)+len(b))
	for key, value := range a {
		result[key] = value
	}
	for key, value := range b {
		result[key] += value
	}
	return result
}
