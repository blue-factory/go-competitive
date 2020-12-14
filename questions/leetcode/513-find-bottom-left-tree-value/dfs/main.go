package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type withHeight struct {
		IsLeft bool
		Height int
		Tn     *TreeNode
	}

	stack := []*withHeight{&withHeight{
		IsLeft: true,
		Height: 0,
		Tn:     root,
	}}

	height := -1
	value := root.Val

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		fmt.Println(node.Tn.Val)
		// fmt.Println(node.Tn.Left)
		// fmt.Println(node.Tn.Right)

		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]

		if node.Height > height {
			height = node.Height
			value = node.Tn.Val
		}

		if node.Tn.Left != nil {
			fmt.Println("JAJAAJ")
			stack = append(stack, &withHeight{
				IsLeft: true,
				Height: node.Height + 1,
				Tn:     node.Tn.Left,
			})
		}

		if node.Tn.Right != nil {
			fmt.Println("JEJEJE")
			stack = append(stack, &withHeight{
				IsLeft: false,
				Height: node.Height + 1,
				Tn:     node.Tn.Right,
			})
		}
	}

	return value
}
