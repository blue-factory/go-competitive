package main

// TreeNode Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	// define withLevel struct used in stack
	type withLevel struct {
		tn    *TreeNode
		level int
	}

	// define avg node to save sum and count number
	type avg struct {
		sum   float64
		count float64
	}

	// create stack
	stack := []withLevel{withLevel{
		tn:    root,
		level: 0,
	}}
	tmp := []avg{avg{
		sum:   float64(root.Val),
		count: 1,
	}}

	// max level
	level := 0

	// iterate over stack
	for len(stack) > 0 {
		// get top value from stack
		node := stack[len(stack)-1]

		// remove current node from stack
		stack = stack[:len(stack)-1]

		// define next level
		nextLevel := node.level + 1
		if nextLevel > level {
			level = nextLevel
		}

		// append position in tmp slice
		tmp = append(tmp, avg{})

		// get left value and push on stack
		if node.tn.Left != nil {
			tmp[nextLevel].sum = tmp[nextLevel].sum + float64(node.tn.Left.Val)
			tmp[nextLevel].count = tmp[nextLevel].count + 1

			stack = append(stack, withLevel{
				tn:    node.tn.Left,
				level: nextLevel,
			})
		}

		// get right value and push on stack
		if node.tn.Right != nil {
			tmp[nextLevel].sum = tmp[nextLevel].sum + float64(node.tn.Right.Val)
			tmp[nextLevel].count = tmp[nextLevel].count + 1

			stack = append(stack, withLevel{
				tn:    node.tn.Right,
				level: nextLevel,
			})
		}
	}

	// iterate tmp and calculate in res avg value for any level
	res := []float64{tmp[0].sum}
	for i := 1; i < level; i++ {
		res = append(res, tmp[i].sum/tmp[i].count)
	}

	return res
}
