package main

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	occurrences := make(map[int]int)

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]

		occurrences[node.Val] = occurrences[node.Val] + 1

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	res := []int{}

	max := MinInt
	for _, sum := range occurrences {
		if sum > max {
			max = sum
		}
	}

	for k, sum := range occurrences {
		if sum == max {
			res = append(res, k)
		}
	}

	return res
}
