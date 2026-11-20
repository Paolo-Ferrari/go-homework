package tasks

func UniqueInPlace(s []int) []int {
	numbers := make(map[int]bool)

	position := 0
	for _, value := range s {
		if !numbers[value] {
			numbers[value] = true

			s[position] = value
			position++
		}
	}
	return s
}
