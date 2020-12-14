package main

func combinationSum(candidates []int, target int) [][]int {
	combinations := [][]int{}
	combinationSumImpl(target, []int{}, &combinations, candidates)
	return combinations
}

func combinationSumImpl(remainder int, combination []int, combinations *[][]int, candidates []int) {
	if remainder == 0 {
		sol := make([]int, len(combination))
		copy(sol, combination)
		(*combinations) = append((*combinations), sol)
		return
	}
	if remainder < 0 {
		return
	}
	for i, c := range candidates {
		combinationSumImpl(remainder-c, append(combination, c), combinations, candidates[i:])
	}
}
