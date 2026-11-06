package tasks

func TrimSet(s, cutset string) string {
	cut := make(map[rune]struct{})
	for _, r := range cutset {
		cut[r] = struct{}{}
	}
	runes := []rune(s)
	start, end := 0, len(runes)
	for start < end {
		if _, ok := cut[runes[start]]; !ok {
			break
		}
		start++
	}
	for start < end {
		if _, ok := cut[runes[end-1]]; !ok {
			break
		}
		end--
	}
	return string(runes[start:end])
}
