package main

// https://leetcode.com/problems/add-digits

import "strconv"

func addDigits(num int) int {
	stack := []int{num}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		n := strconv.Itoa(top)
		if len(n) < 2 {
			return top
		}

		sum := 0
		for _, nn := range n {
			value, _ := strconv.Atoi(string(nn))
			sum += value
		}

		if sum < 10 {
			return sum
		}

		stack = append(stack, sum)
	}

	return 0
}
