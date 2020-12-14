package main

// https://leetcode.com/problems/single-number

func singleNumber(nums []int) int {
	b := 0

	for _, n := range nums {
		b ^= n
	}

	return b
}
