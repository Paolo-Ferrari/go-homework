package tasks

func MostFrequent(nums []int) (value int, count int, ok bool) {
	frequencies := make(map[int]int)
	for _, n := range nums {
		frequencies[n]++
	}
	for n, frequency := range frequencies {
		if !ok || frequency > count || (frequency == count && n < value) {
			value, count, ok = n, frequency, true
		}
	}
	return
}
