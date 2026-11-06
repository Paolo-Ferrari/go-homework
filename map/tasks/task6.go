package tasks

func SumValues(m map[string]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}
