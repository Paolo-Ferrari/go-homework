package main

import (
	"fmt"

	"github.com/Paolo-Ferrari/go-homework/slice/tasks"
)

func main() {
	fmt.Println("Homework: Slice")

	// task1
	fmt.Println("task1:", tasks.Sum([]int{1, 2, 3, 4, 5}))

	// task2
	min, ok := tasks.Min([]int{5, 2, 8, 1, 9})
	fmt.Println("task2:", min, ok)

	// task3
	max, ok := tasks.Max([]int{5, 2, 8, 1, 9})
	fmt.Println("task3:", max, ok)

	// task4
	fmt.Println("task4:", tasks.CountEven([]int{1, 2, 3, 4, 5, 6}))

	// task5
	fmt.Println("task5:", tasks.Contains([]int{1, 2, 3}, 2))

	// task6
	fmt.Println("task6:", tasks.CountOccurrences([]int{1, 2, 2, 3, 2}, 2))

	// task7
	fmt.Println("task7:", tasks.IndexOf([]int{5, 2, 8, 2}, 2))

	// task8
	fmt.Println("task8:", tasks.IndexesOf([]int{5, 2, 8, 2, 9, 2}, 2))

	// task9
	s9 := []int{1, 2, 3, 2, 4}
	tasks.ReplaceAll(s9, 2, 9)
	fmt.Println("task9:", s9)

	// task10
	s10 := []int{1, 2, 3, 4, 5}
	tasks.Reverse(s10)
	fmt.Println("task10:", s10)

	// task11
	s11 := []int{1, 2, 3, 4, 5}
	tasks.SwapPairs(s11)
	fmt.Println("task11:", s11)

	// task12
	s12 := []int{1, 2, 3, 4, 5}
	err := tasks.ReverseRange(s12, 1, 3)
	fmt.Println("task12:", s12, err)

	// task13
	s13 := []int{1, 2, 3, 4, 5}
	result13, err := tasks.RemoveAt(s13, 2)
	fmt.Println("task13:", result13, err)

	// task14
	s14 := []int{1, 2, 3, 4, 5}
	result14, err := tasks.RemoveAtFast(s14, 1)
	fmt.Println("task14:", result14, err)

	// task15
	s15 := []int{1, 2, 3, 2, 4, 2}
	result15 := tasks.RemoveAll(s15, 2)
	fmt.Println("task15:", result15)

	// task16
	s16 := []int{1, 2, 3, 2, 4, 2}
	result16 := tasks.RemoveAllInPlace(s16, 2)
	fmt.Println("task16:", result16)

	// task17
	s17 := []int{1, 2, 3, 4, 5, 6}
	result17 := tasks.KeepEvenInPlace(s17)
	fmt.Println("task17:", result17)

	// task18
	s18 := []int{1, 0, 2, 0, 3, 0, 4}
	tasks.MoveZerosToEnd(s18)
	fmt.Println("task18:", s18)

	// task19
	s19 := []int{1, 0, 2, 0, 3, 0, 4}
	result19 := tasks.CompactNonZero(s19)
	fmt.Println("task19:", result19)

	// task20
	s20 := []int{1, 2, 3, 4, 5}
	result20, err := tasks.InsertAt(s20, 2, 99)
	fmt.Println("task20:", result20, err)

	// task21
	s21 := []int{10, 20, 30, 40, 50}
	values21 := []int{11, 69}
	resule21, err := tasks.InsertMany(s21, 3, values21)
	fmt.Println("task21:", resule21, err)

	// task22
	s22 := []int{11, 22, 33, 44, 55}
	result22, err := tasks.DeleteRange(s22, 1, 3)
	fmt.Println("task22:", result22, err)

	//task23
	s23 := []int{11, 22, 33, 44, 55}
	left23 := 2
	right23 := 4
	values23 := []int{99, 69}
	result23, err := tasks.ReplaceRange(s23, left23, right23, values23)
	fmt.Println("tasks23:", result23, err)

	//task24
	s24 := []int{11, 22, 33, 44, 55}
	values24 := []int{69, 79, 99}
	result24, err := tasks.Splice(s24, 1, 3, values24)
	fmt.Println("task24:", result24, err)

	//task25
	s25 := []int{11, 22, 33}
	result25 := tasks.Clone(s25)
	fmt.Println("task25:", result25)

	//task26
	s26 := []int{11, 22, 33, 44, 55}
	result26, err := tasks.CloneRange(s26, 1, 3)
	fmt.Println("task26:", result26, err)

	//task27
	a27 := []int{11, 22, 33}
	b27 := []int{44, 55}
	result27 := tasks.Concat(a27, b27)
	fmt.Println("task27:", result27)

	//task28
	a28 := []int{11, 22, 33}
	b28 := []int{44, 55}
	c28 := []int{99, 69}
	result28 := tasks.ConcatMany(a28, b28, c28)
	fmt.Println("task28:", result28)

	//task29
	s29 := []int{11, 22, 11, 33, 44, 44, 44, 55, 55, 66}
	result29 := tasks.Unique(s29)
	fmt.Println("task29:", result29)

	//task30
	s30 := []int{11, 22, 33, 11, 33, 22, 66, 44, 44, 55}
	result30 := tasks.UniqueInPlace(s30)
	fmt.Println("task30:", result30)
}
