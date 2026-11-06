package tasks

func Clone(m map[string][]int) map[string][]int {
	if m == nil {
		return nil
	}
	result := make(map[string][]int, len(m))
	for key, values := range m {
		result[key] = append([]int(nil), values...)
	}
	return result
}
