package tasks

func IndexesOf(s []int, target int) []int {
	indexes := []int{}

	for i, num := range s {
		if num == target {
			indexes = append(indexes, i)
		}
	}

	return indexes
}
