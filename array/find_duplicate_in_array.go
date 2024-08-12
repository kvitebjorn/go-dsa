package array

// FindDuplicate solves the problem in O(n) time and O(1) space.
/*
func FindDuplicate(list []int) int {
	for _, item := range list {
		itemIndex := item - 1
		if list[itemIndex] < 0 {
			return list[itemIndex] * -1
		}
		list[itemIndex] *= -1
	}
	return -1
}
*/

// 1st iteration:
// Oops, didn't realize that `list` is guaranteed to be the sequence 1,2..n :P
// in other words, a map is not required, and violated the O(1) space req anyway...
/*
func FindDuplicate(list []int) int {
	M := make(map[int]bool)
	for _, elem := range list {
		_, ok := M[elem]
		if !ok {
			M[elem] = true
		} else {
			return elem
		}
	}
	return -1
}
*/

func FindDuplicate(list []int) int {
	for idx, elem := range list {
		if idx+1 != elem {
			return elem
		}
	}
	return -1
}
