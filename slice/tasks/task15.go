package tasks

func RemoveAll(s []int, target int) []int {
	result := make([]int, 0)

	for _, num := range s {
		if num != target {
			result = append(result, num)
		}
	}

	return result
}
