package tasks

func Min(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	min := s[0]

	for _, num := range s {
		if num < min {
			min = num
		}
	}

	return min, true
}
