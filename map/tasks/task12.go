package tasks

func Difference(a, b []int) []int {
	inB := make(map[int]struct{}, len(b))
	for _, n := range b {
		inB[n] = struct{}{}
	}
	seen := make(map[int]struct{})
	result := make([]int, 0)
	for _, n := range a {
		if _, exists := inB[n]; exists {
			continue
		}
		if _, exists := seen[n]; exists {
			continue
		}
		seen[n] = struct{}{}
		result = append(result, n)
	}
	return result
}
