package tasks

import "fmt"

func Invert(m map[string]int) (map[int]string, error) {
	result := make(map[int]string, len(m))
	for key, value := range m {
		if previous, exists := result[value]; exists {
			return nil, fmt.Errorf("duplicate value %d for keys %q and %q", value, previous, key)
		}
		result[value] = key
	}
	return result, nil
}
