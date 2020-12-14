package main

// https://leetcode.com/problems/two-sum

func twoSum(nums []int, target int) []int {
	nn := make(map[int]int)

	for i, n := range nums {
		key := target - n
		v, ok := nn[key]
		if ok {
			return []int{v, i}
		}

		nn[n] = i
	}

	return []int{}
}
