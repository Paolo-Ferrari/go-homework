package tasks

import "errors"

func ReverseRange(s []int, left, right int) error {
	if left < 0 || right >= len(s) || left > right {
		return errors.New("invalid range")
	}

	for left < right {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}

	return nil
}
