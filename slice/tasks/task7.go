package tasks

func IndexOf(s []int, target int) int {
	for i, num := range s {
		if num == target {
			return i
		}
	}

	return -1
}
