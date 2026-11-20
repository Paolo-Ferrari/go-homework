// Прикольная задача минут 10 думал над страчками 11-14
package tasks

func ConcatMany(parts ...[]int) []int {
	length := 0
	for _, part := range parts {
		length += len(part)
	}

	result := make([]int, length)

	copyHelper := 0
	for _, part := range parts {
		copyHelper += copy(result[copyHelper:], part)
	}

	return result
}
