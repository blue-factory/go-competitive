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

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type withValues struct {
		TN     *TreeNode
		Height int
	}

	stack := []*withValues{&withValues{root, 0}}
	sum := make(map[int]int)
	maxHeight := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]

		if node.TN.Left == nil && node.TN.Right == nil {
			sum[node.Height] = sum[node.Height] + node.TN.Val
			if node.Height > maxHeight {
				maxHeight = node.Height
			}
		}

		if node.TN.Left != nil {
			stack = append(stack, &withValues{node.TN.Left, node.Height + 1})
		}

		if node.TN.Right != nil {
			stack = append(stack, &withValues{node.TN.Right, node.Height + 1})
		}
	}

	return sum[maxHeight]
}
