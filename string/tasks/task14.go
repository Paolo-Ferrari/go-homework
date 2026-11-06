package tasks

import "fmt"

func CountOverlapping(s, sub string) (int, error) {
	if sub == "" {
		return 0, fmt.Errorf("substring must not be empty")
	}
	count := 0
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			count++
		}
	}
	return count, nil
}
