package tasks

func Unique(nums []int) []int {
	seen := make(map[int]struct{})
	result := make([]int, 0, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			continue
		}
		seen[n] = struct{}{}
		result = append(result, n)
	}
	return result
}
