// После task24 эта задача - издевательство,
// я 24 на протяжении 1.5-2 часов решал,
// пока искал про то как работает splice, пока всё остальное...
package tasks

func Clone(s []int) []int {
	result := make([]int, len(s))
	copy(result, s)
	return result
}
