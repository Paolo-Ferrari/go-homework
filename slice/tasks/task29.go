// Пмогла задача с урока, и правило если не занешь
// как решать то тупо закинь в мапу
// Здесь же в комментарии (под этим решением) дуркая реализация,
// на всякий, если я не правильно понял условие

package tasks

func Unique(s []int) []int {
	numbers := make(map[int]bool)

	result := make([]int, 0, len(s))

	for _, value := range s {
		if !numbers[value] {
			numbers[value] = true
			result = append(result, value)
		}
	}

	return result
}

// func Unique(s []int) []int {

// 	counter := make(map[int]int)
// 	for _, value := range s {
// 		counter[value]++
// 	}

// 	result := make([]int, 0, len(s))

// 	for _, value := range s {
// 		if counter[value] == 1 {
// 			result = append(result, value)
// 		}
// 	}

// 	return result
// }
