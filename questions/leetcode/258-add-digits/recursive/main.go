package main

// https://leetcode.com/problems/add-digits

import "strconv"

func addDigits(num int) int {
	n := strconv.Itoa(num)
	if len(n) < 2 {
		return num
	}

	sum := 0
	for _, nn := range n {
		value, _ := strconv.Atoi(string(nn))
		sum += value
	}

	if sum < 10 {
		return sum
	}

	return addDigits(sum)
}
