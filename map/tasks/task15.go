package tasks

import "unicode"

func UniqueFold(words []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(words))
	for _, word := range words {
		keyRunes := []rune(word)
		for i, r := range keyRunes {
			keyRunes[i] = unicode.ToLower(r)
		}
		key := string(keyRunes)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, word)
	}
	return result
}
