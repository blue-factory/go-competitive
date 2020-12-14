package main

const MaxInt = int(^uint(0) >> 1)
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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type withData struct {
		TN  *TreeNode
		Min int
		Max int
	}

	stack := []*withData{&withData{root, MinInt, MaxInt}}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]

		if node.TN.Val <= node.Min || node.TN.Val >= node.Max {
			return false
		}

		if node.TN.Left != nil {
			stack = append(stack, &withData{node.TN.Left, node.Min, node.TN.Val})
		}

		if node.TN.Right != nil {
			stack = append(stack, &withData{node.TN.Right, node.TN.Val, node.Max})
		}
	}

	return true
}
