package tasks

import "sort"

func SymmetricDifference(a, b []int) []int {
	setA := make(map[int]struct{}, len(a))
	setB := make(map[int]struct{}, len(b))
	for _, n := range a {
		setA[n] = struct{}{}
	}
	for _, n := range b {
		setB[n] = struct{}{}
	}

	result := make([]int, 0)
	for n := range setA {
		if _, exists := setB[n]; !exists {
			result = append(result, n)
		}
	}
	for n := range setB {
		if _, exists := setA[n]; !exists {
			result = append(result, n)
		}
	}
	sort.Ints(result)
	return result
}
