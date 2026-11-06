package tasks

func LongestCommonPrefix(items []string) string {
	if len(items) == 0 {
		return ""
	}
	prefix := []rune(items[0])
	for _, item := range items[1:] {
		runes := []rune(item)
		i := 0
		for i < len(prefix) && i < len(runes) && prefix[i] == runes[i] {
			i++
		}
		prefix = prefix[:i]
		if len(prefix) == 0 {
			return ""
		}
	}
	return string(prefix)
}
