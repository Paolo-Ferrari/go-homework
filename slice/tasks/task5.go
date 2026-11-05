package tasks

func Contains(s []int, target int) bool {
	for _, num := range s {
		if num == target {
			return true
		}
	}

	return false
}
