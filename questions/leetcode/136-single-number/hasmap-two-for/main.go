package main

func singleNumber(nums []int) int {
	values := make(map[int]bool)

	for _, n := range nums {
		_, ok := values[n]
		if ok {
			values[n] = true
		} else {
			values[n] = false
		}
	}

	for k, v := range values {
		if v == false {
			return k
		}
	}

	return -1
}
