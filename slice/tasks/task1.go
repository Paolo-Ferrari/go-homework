package tasks

func Sum(s []int) int {
	sum := 0

	for _, num := range s {
		sum += num
	}

	return sum
}
