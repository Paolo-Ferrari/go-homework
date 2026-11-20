package tasks

import "errors"

func CloneRange(s []int, left, right int) ([]int, error) {
	if left < 0 || left > len(s) || right < 0 || right > len(s) || left > right {
		return nil, errors.New("invalid range")
	}

	result := make([]int, right-left+1)
	copy(result, s[left:right+1])

	return result, nil
}
