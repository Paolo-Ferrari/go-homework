package tasks

func DeleteZeroValues(m map[string]int) {
	for key, value := range m {
		if value == 0 {
			delete(m, key)
		}
	}
}
