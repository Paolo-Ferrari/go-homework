package tasks

import "errors"

func ReplaceRange(s []int, left, right int, values []int) ([]int, error) {
	if left < 0 || right < 0 || left > len(s) || right > len(s) || left > right {
		return nil, errors.New("invalid range")
	}

	result := make([]int, 0, len(s)-(right-left)+len(values))
	result = append(result, s[:left]...)
	result = append(result, values...)
	result = append(result, s[right:]...)

	return result, nil
}
