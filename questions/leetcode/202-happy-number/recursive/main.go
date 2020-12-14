package main

import "strconv"

// https://leetcode.com/problems/happy-number

func isHappy(n int) bool {
	visited := make(map[int]bool)

	return isHappyImpl(n, visited)
}

func isHappyImpl(n int, visited map[int]bool) bool {
	numStr := strconv.Itoa(n)

	sum := 0
	for _, c := range numStr {
		num, _ := strconv.Atoi(string(c))
		sum += num * num
	}

	if sum == 1 {
		return true
	}

	if _, exists := visited[sum]; !exists {
		visited[sum] = true
	} else {
		return false
	}

	return isHappyImpl(sum, visited)
}
