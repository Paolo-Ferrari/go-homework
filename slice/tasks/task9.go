package tasks

func ReplaceAll(s []int, old, new int) {
	for i := range s {
		if s[i] == old {
			s[i] = new
		}
	}
}
