package main

// https://leetcode.com/problems/happy-number

import "strconv"

func isHappy(n int) bool {
	stack := []string{strconv.Itoa(n)}
	visited := make(map[int]bool)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		sum := 0
		for _, c := range top {
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

		stack = append(stack, strconv.Itoa(sum))
	}

	return false
}
