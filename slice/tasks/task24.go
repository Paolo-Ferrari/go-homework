package tasks

import "errors"

func Splice(s []int, start, deleteCount int, values []int) ([]int, error) {
	if start < 0 || deleteCount+start > len(s) || start > len(s) || deleteCount < 0 {
		return nil, errors.New("invalid range")
	}

	if (len(s) - deleteCount + len(values)) > cap(s) {
		result := make([]int, len(s)-deleteCount+len(values))

		copy(result, s[:start])
		copy(result[start:], values)
		copy(result[start+len(values):], s[start+deleteCount:])

		return result, nil
	}

	copy(s[start:], values)
	copy(s[start+len(values):], s[start+deleteCount:])
	s = s[:(len(s) - deleteCount + len(values))]

	return s, nil
}
