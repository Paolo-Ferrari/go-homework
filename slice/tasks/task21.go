package tasks

import "errors"

func InsertMany(s []int, i int, values []int) ([]int, error) {
	if i < 0 || i > len(s) {
		return nil, errors.New("invalid index")
	}

	result := make([]int, 0, len(s)+len(values))
	result = append(result, s[:i]...)
	result = append(result, values...)
	result = append(result, s[i:]...)

	return result, nil
}
