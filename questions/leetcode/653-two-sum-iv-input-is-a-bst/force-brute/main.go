package main

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

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{root}
	nums := []int{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]

		nums = append(nums, node.Val)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	if len(nums) == 1 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == k {
				return true
			}
		}
	}

	return false
}
