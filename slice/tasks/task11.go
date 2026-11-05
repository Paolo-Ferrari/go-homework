package tasks

func SwapPairs(s []int) {
	for i := 0; i+1 < len(s); i += 2 {
		s[i], s[i+1] = s[i+1], s[i]
	}
}
