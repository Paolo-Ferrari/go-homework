//Мало того, что задачу решить надо, хотя и не сложную,
//так ещё и функцию придумать, для примера

package tasks

func MapValues(m map[string]int, f func(int) int) map[string]int {
	result := make(map[string]int)

	for key, value := range m {
		result[key] = f(value)
	}

	return result
}
