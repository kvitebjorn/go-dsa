package array

// ProductOfAllOtherElements solves the problem in O(n) time and O(1) space.
/*
func ProductOfAllOtherElements(list []int) []int {
	if len(list) == 0 {
		return list
	}

	right := make([]int, len(list))

	product := 1
	for i := len(list) - 1; i >= 0; i-- {
		product *= list[i]
		right[i] = product
	}

	product = list[0]
	for i := 0; i < len(list); i++ {
		if i == 0 {
			list[i] = right[1]
			continue
		}
		value := list[i]

		if i != len(list)-1 {
			list[i] = right[i+1] * product
		} else {
			list[i] = product
		}

		product *= value
	}

	return list
}

func ProductOfAllOtherElements(list []int) []int {
	B := make([]int, len(list))

	for idx := range B {
		B[idx] = 1
	}

	// this is just syntactic sugar for a nested loop, i have to figure something else out
	aIdx := 0
	bIdx := 0
	for bIdx < len(list) {
		if aIdx > len(list)-1 {
			aIdx = 0
			bIdx = bIdx + 1
			continue
		}
		if aIdx == bIdx {
			aIdx = aIdx + 1
			continue
		}

		B[bIdx] = B[bIdx] * list[aIdx]
		aIdx = aIdx + 1
	}

	return B
}
*/

func ProductOfAllOtherElements(list []int) []int {
	if len(list) == 0 {
		return list
	}

	B := make([]int, len(list))

	// accumulate the product from the left side
	// [1 2 3 4] -> [1 2 6 24]
	product := 1
	for i := 0; i < len(list); i++ {
		product *= list[i]
		B[i] = product
	}

	// now go from the right side
	// [1 2 6 24] -> [(2 * 3 * 4) (1 * 3 * 4) (2 * 4) 6]
	product = list[len(list)-1]
	for i := len(list) - 1; i >= 0; i-- {
		if i == len(list)-1 {
			list[i] = B[i-1]
			continue
		}
		value := list[i]

		if i > 0 {
			list[i] = B[i-1] * product // multiply the product on the left with the product to the right
		} else {
			list[i] = product
		}

		product *= value // keep the running product going from right to left
	}

	return list
}
