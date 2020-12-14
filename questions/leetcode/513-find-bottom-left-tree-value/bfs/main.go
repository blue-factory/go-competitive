package main

import "container/list"

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

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		node := queue.Front()
		nodeValue := node.Value.(*TreeNode)

		root = nodeValue

		if nodeValue.Right != nil {
			queue.PushBack(nodeValue.Right)
		}

		if nodeValue.Left != nil {
			queue.PushBack(nodeValue.Left)
		}

		queue.Remove(node)
	}

	return root.Val
}
