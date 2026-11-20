package tasks

import "errors"

func DeleteRange(s []int, left, right int) ([]int, error) {
	if left < 0 || right < 0 || left > len(s) || right > len(s) {
		return nil, errors.New("invalid rane")
	}

	if left == right {
		result := make([]int, 0, 1)
		return append(result, s[right]), nil
	}

	result := make([]int, 0, len(s)-(right-left))
	result = append(result, s[:left]...)
	result = append(result, s[right:]...)

	return result, nil
}
