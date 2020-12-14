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
	nums := make(map[int]bool)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]

		key := k - node.Val
		if nums[key] {
			return true
		}

		nums[node.Val] = true

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return false
}
