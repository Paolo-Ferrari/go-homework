package tasks

func CountInts(nums []int) map[int]int {
	result := make(map[int]int)
	for _, n := range nums {
		result[n]++
	}
	return result
}
