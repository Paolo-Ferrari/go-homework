package tasks

func CompactNonZero(s []int) []int {
	write := 0

	for _, num := range s {
		if num != 0 {
			s[write] = num
			write++
		}
	}

	return s[:write]
}
