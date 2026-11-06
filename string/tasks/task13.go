package tasks

import "fmt"

func CountOccurrences(s, sub string) (int, error) {
	if sub == "" {
		return 0, fmt.Errorf("substring must not be empty")
	}
	count := 0
	for i := 0; i <= len(s)-len(sub); {
		if s[i:i+len(sub)] == sub {
			count++
			i += len(sub)
		} else {
			i++
		}
	}
	return count, nil
}
