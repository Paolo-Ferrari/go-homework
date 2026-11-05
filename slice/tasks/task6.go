package tasks

func CountOccurrences(s []int, target int) int {
	count := 0

	for _, num := range s {
		if num == target {
			count++
		}
	}

	return count
}
