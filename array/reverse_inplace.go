package array

// ReverseInPlace solves the problem in O(n) time and O(1) space.
/*
func ReverseInPlace(list []int, start, end int) {
	for i := start; i <= start+end/2 && i < end-i+start; i++ {
		list[i], list[end-i+start] = list[end-i+start], list[i]
	}
}
*/

func ReverseInPlace(list []int, start, end int) {
	for i, j := start, end; i < j; i++ {
		list[i], list[j] = list[j], list[i]
		j--
	}
}
