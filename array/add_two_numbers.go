package array

import "strconv"

// AddTwoNumbers solves the problem in O(n) time and O(1) space.
/*
func AddTwoNumbers(num1, num2 []int) []int {
	num1, num2 = equalizeLengths(num1, num2)
	carry := false
	for i := len(num1) - 1; i > -1; i-- {
		num1[i] += num2[i]
		if carry {
			num1[i]++
			carry = false
		}
		if num1[i] >= 10 {
			num1[i] -= 10
			carry = true
		}
	}
	if carry {
		num1 = append([]int{1}, num1...)
	}
	return num1
}

func equalizeLengths(num1, num2 []int) ([]int, []int) {
	diff := int(math.Abs(float64(len(num2) - len(num1))))
	zeros := make([]int, diff)
	if len(num2) > len(num1) {
		num1 = append(zeros, num1...)
	} else if len(num1) > len(num2) {
		num2 = append(zeros, num2...)
	}
	return num1, num2
}
*/

func AddTwoNumbers(num1, num2 []int) []int {
	acc1 := ""
	acc2 := ""

	for _, e := range num1 {
		acc1 += strconv.Itoa(e)
	}
	for _, e := range num2 {
		acc2 += strconv.Itoa(e)
	}

	n1, _ := strconv.Atoi(acc1)
	n2, _ := strconv.Atoi(acc2)

	res := n1 + n2
	resStr := strconv.Itoa(res)

	resArr := []int{}
	for _, e := range resStr {
		r := int(e - '0')
		resArr = append(resArr, r)
	}

	return resArr
}
