package tasks

func Max(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	max := s[0]

	for _, num := range s {
		if num > max {
			max = num
		}
	}

	return max, true
}
