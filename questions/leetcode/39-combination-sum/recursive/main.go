package main

// https://leetcode.com/problems/combination-sum

func combinationSum(candidates []int, target int) [][]int {
	results := [][]int{}

	combinationSum2Impl(0, []int{}, candidates, target, &results)

	return results
}

func combinationSum2Impl(sum int, candidate []int, candidates []int, target int, results *[][]int) {
	if sum == target {
		sol := make([]int, len(candidate))
		copy(sol, candidate)

		*results = append(*results, sol)
		return
	}

	if sum > target {
		return
	}

	for i, c := range candidates {
		combinationSum2Impl(sum+c, append(candidate, c), candidates[i:], target, results)
	}
}
