// Пример для этой задачи придумать сложнее,
// чем решение для неё, я серьёзно на пример, в main.go
// потратил больше времени
package tasks

func FileterMap(m map[string]int, keep func(string, int) bool) map[string]int {
	result := make(map[string]int)

	for key, value := range m {
		if keep(key, value) {
			result[key] = value
		}
	}

	return result
}
