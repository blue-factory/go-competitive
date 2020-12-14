package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	tmp := []*TreeNode{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		tmp = append(tmp, node)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	for i, n := range tmp {
		n.Left = nil

		if i == len(tmp)-1 {
			n.Right = nil
		} else {
			n.Right = tmp[i+1]
		}
	}
}
