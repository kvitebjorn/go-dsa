package array

import (
	"sort"
)

// ZeroSumTriplets solves the problem in O(n^2) time and O(1) space.
/*
func ZeroSumTriplets(list []int) [][]int {
	output := make([][]int, 0)
	if len(list) < 3 {
		return output
	}

	sort.Ints(list)
	for i, n := range list {
		if i > 0 && n == list[i-1] {
			continue
		}

		l, r := i+1, len(list)-1
		for l < r {
			threeSum := n + list[l] + list[r]
			if threeSum > 0 {
				r--
				continue
			}
			if threeSum < 0 {
				l++
				continue
			}
			output = append(output, []int{n, list[l], list[r]})
			l++
			for list[l] == list[l-1] && l < r {
				l++
			}
		}
	}
	return output
}
*/

func ZeroSumTriplets(list []int) [][]int {
	T := [][]int{}
	if len(list) < 3 {
		return T
	}
	sort.Ints(list)
	for idx, n := range list {
		if n > 0 {
			// we need to start with at least one negative number or zeros
			break
		}

		if idx > 0 && list[idx-1] == n {
			// skip duplicates to satisfy the `unique` constraint
			continue
		}

		// slider approach from both sides works nicely due to the sort
		// if the sum is < 0, we know we must increment the left side (too small)
		// if the sum is > 0, we know we must decrement the right side (too high)
		l, r := idx+1, len(list)-1
		for l < r {
			threeSum := n + list[l] + list[r]
			if threeSum < 0 {
				l++
				continue
			}
			if threeSum > 0 {
				r--
				continue
			}
			T = append(T, []int{n, list[l], list[r]})
			l++
			for list[l-1] == list[l] && l < r {
				// skip duplicates to satisfy the `unique` constraint
				l++
			}
		}
	}
	return T
}
