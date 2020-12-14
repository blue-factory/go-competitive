package main

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func indexOf(s []int, e int) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

func singleNumber(nums []int) int {
	n := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if contains(n, nums[i]) {
			index := indexOf(n, nums[i])
			n = append(n[:index], n[index+1:]...)
		} else {
			n = append(n, nums[i])
		}
	}

	return n[0]
}
