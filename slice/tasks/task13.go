package tasks

import "errors"

func RemoveAt(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return nil, errors.New("invalid index")
	}

	copy(s[i:], s[i+1:])

	return s[:len(s)-1], nil
}
