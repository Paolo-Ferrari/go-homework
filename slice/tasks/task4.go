package tasks

func CountEven(s []int) int {
	count := 0

	for _, num := range s {
		if num%2 == 0 {
			count++
		}
	}

	return count
}
