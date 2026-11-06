package tasks

func ContainsAll(have, need []string) bool {
	set := make(map[string]struct{}, len(have))
	for _, value := range have {
		set[value] = struct{}{}
	}
	for _, value := range need {
		if _, ok := set[value]; !ok {
			return false
		}
	}
	return true
}
