package main

import (
	"fmt"

	maptasks "github.com/Paolo-Ferrari/go-homework/map/tasks"
)

func main() {
	fmt.Println("Homework: Map")
	fmt.Println("Task 1:", maptasks.CountInts([]int{1, 2, 1, 3, 2, 1}))
	fmt.Println("Task 2:", maptasks.LastIndex([]string{"css", "js", "go"}))
	fmt.Println("Task 3:", maptasks.Unique([]int{3, 1, 3, 2, 1}))
	fmt.Println("Task 4:", maptasks.ContainsAll([]string{"go", "map"}, []string{"map"}))
	fmt.Println("Task 5:", maptasks.SortedKeys(map[string]int{"z": 1, "a": 2}))
	fmt.Println("Task 6:", maptasks.SumValues(map[string]int{"a": 2, "b": 3}))
	inverted, err := maptasks.Invert(map[string]int{"a": 1, "b": 2})
	fmt.Println("Task 7:", inverted, err)
	fmt.Println("Task 8:", maptasks.MergeCounts(map[string]int{"a": 2}, map[string]int{"a": 3, "b": 1}))
	counts := map[string]int{"a": 0, "b": 2}
	maptasks.DeleteZeroValues(counts)
	fmt.Println("Task 9:", counts)
	fmt.Println("Task 10:", maptasks.Clone(map[string][]int{"a": {1, 2}}))
	fmt.Println("Task 11:", maptasks.Equal(map[string]int{"a": 1}, map[string]int{"a": 1}))
	fmt.Println("Task 12:", maptasks.Difference([]int{1, 2, 2, 3}, []int{2}))
	fmt.Println("Task 13:", maptasks.Intersection([]string{"a", "b", "a"}, []string{"b", "a"}))
	fmt.Println("Task 14:", maptasks.SymmetricDifference([]int{1, 2, 3}, []int{3, 4}))
	fmt.Println("Task 15:", maptasks.UniqueFold([]string{"Go", "go", "GOLANG"}))
	value, count, ok := maptasks.MostFrequent([]int{4, 2, 4, 2, 2})
	fmt.Println("Task 16:", value, count, ok)
	fmt.Println("Task 17:", stringResult(maptasks.FirstUniqueRune("aabbc")))
	fmt.Println("Task 18:", maptasks.AreAnagrams("Dormitory", "dirty room"))
	fmt.Println("Task 19:", maptasks.WordFrequency("Go go, язык 123!"))
	fmt.Println("Task 20:", maptasks.TopKWords("go map go string map go", 2))
}

func stringResult(r rune, ok bool) string {
	if !ok {
		return "<none>"
	}
	return string(r)
}
