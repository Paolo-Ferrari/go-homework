package tasks

func GoupByParity(nums []int) map[string][]int {
	result := make(map[string][]int)

	for _, num := range nums {
		if num%2 == 0 {
			result["even"] = append(result["even"], num)
		} else {
			result["odd"] = append(result["odd"], num)
		}
	}

	return result
}
