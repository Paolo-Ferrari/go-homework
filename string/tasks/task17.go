package tasks

import "fmt"

func Split(s, sep string) ([]string, error) {
	if sep == "" {
		return nil, fmt.Errorf("separator must not be empty")
	}
	result := make([]string, 0)
	start := 0
	for i := 0; i <= len(s)-len(sep); {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}
	return append(result, s[start:]), nil
}
