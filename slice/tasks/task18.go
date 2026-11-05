package tasks

func MoveZerosToEnd(s []int) {
	write := 0

	for _, num := range s {
		if num != 0 {
			s[write] = num
			write++
		}
	}

	for i := write; i < len(s); i++ {
		s[i] = 0
	}
}
