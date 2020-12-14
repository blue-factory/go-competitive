package main

func twoSum(numbers []int, target int) []int {
	nums := make(map[int]int)

	for index, n := range numbers {
		key := target - n
		value, ok := nums[key]
		if ok {
			return []int{value + 1, index + 1}
		}

		nums[n] = index
	}

	return []int{}
}
