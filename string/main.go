package main

import (
	"fmt"

	stringtasks "github.com/Paolo-Ferrari/go-homework/string/tasks"
)

func main() {
	fmt.Println("Homework: String")
	byteCount, runeCount := stringtasks.StringSize("Привет 😀")
	fmt.Println("Task 1:", byteCount, runeCount)
	fmt.Println("Task 2:", stringtasks.IsASCII("Hello!"))
	fmt.Println("Task 3:", stringResult(stringtasks.FirstRune("Go")))
	fmt.Println("Task 4:", stringResult(stringtasks.LastRune("Go")))
	fmt.Println("Task 5:", stringtasks.Reverse("Go 😀"))
	fmt.Println("Task 6:", stringtasks.RemoveRune("banana", 'a'))
	fmt.Println("Task 7:", stringtasks.ReplaceRune("banana", 'a', 'o'))
	fmt.Println("Task 8:", stringtasks.IsPalindrome("казак"))
	fmt.Println("Task 9:", stringtasks.TrimUnicodeSpace("  Go  "))
	fmt.Println("Task 10:", stringtasks.NormalizeSpaces(" Go   is\\t fun "))
	fmt.Println("Task 11:", stringtasks.Contains("hello", "ell"))
	fmt.Println("Task 12:", stringtasks.IndexOf("hello", "ll"))
	occurrences, err := stringtasks.CountOccurrences("aaaa", "aa")
	fmt.Println("Task 13:", occurrences, err)
	overlapping, err := stringtasks.CountOverlapping("aaaa", "aa")
	fmt.Println("Task 14:", overlapping, err)
	fmt.Println("Task 15:", stringtasks.CommonPrefix("привет", "приход"))
	fmt.Println("Task 16:", stringtasks.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	parts, err := stringtasks.Split("a,,b,", ",")
	fmt.Println("Task 17:", parts, err)
	fmt.Println("Task 18:", stringtasks.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println("Task 19:", stringtasks.TrimSet("...hello!!", ".!"))
	fmt.Println("Task 20:", stringtasks.CollapseRuns("аааббв"))
}

func stringResult(r rune, ok bool) string {
	if !ok {
		return "<none>"
	}
	return string(r)
}
