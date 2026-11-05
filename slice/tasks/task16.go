package tasks

func RemoveAllInPlace(s []int, target int) []int {
	write := 0

	for _, num := range s {
		if num != target {
			s[write] = num
			write++
		}
	}

	return s[:write]
}
