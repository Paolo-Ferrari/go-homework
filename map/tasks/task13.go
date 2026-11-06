package tasks

func Intersection(a, b []string) []string {
	inB := make(map[string]struct{}, len(b))
	for _, value := range b {
		inB[value] = struct{}{}
	}
	seen := make(map[string]struct{})
	result := make([]string, 0)
	for _, value := range a {
		if _, ok := inB[value]; !ok {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}
