package main

import (
	"container/list"
	"fmt"
)

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

	// define withLevel struct used in queue
	type withLevel struct {
		tn    *TreeNode
		level int
	}

	// define avg node to save sum and count
	type avg struct {
		sum   float64
		count float64
	}

	// create queue
	queue := list.New()

	// set root level
	queue.PushBack(withLevel{
		tn:    root,
		level: 0,
	})
	tmp := []avg{avg{
		sum:   float64(root.Val),
		count: 1,
	}}

	// max level
	level := 0

	// iterate over queue
	for queue.Len() > 0 {
		// get top queue value
		node := queue.Front()

		// cast node to withLevel
		nodeWithLevel := node.Value.(withLevel)

		fmt.Println(nodeWithLevel.tn.Val, nodeWithLevel.level)

		// fmt.Println("cap tmp:", cap(tmp))

		// define next level
		nextLevel := nodeWithLevel.level + 1
		if nextLevel > level {
			level = nextLevel
		}

		// append position in tmp slice
		tmp = append(tmp, avg{})

		// get left value and push back on queue
		if nodeWithLevel.tn.Left != nil {
			tmp[nextLevel].sum = tmp[nextLevel].sum + float64(nodeWithLevel.tn.Left.Val)
			tmp[nextLevel].count = tmp[nextLevel].count + 1

			queue.PushBack(withLevel{
				tn:    nodeWithLevel.tn.Left,
				level: nextLevel,
			})
		}

		// get right value and push back on queue
		if nodeWithLevel.tn.Right != nil {
			tmp[nextLevel].sum = tmp[nextLevel].sum + float64(nodeWithLevel.tn.Right.Val)
			tmp[nextLevel].count = tmp[nextLevel].count + 1

			queue.PushBack(withLevel{
				tn:    nodeWithLevel.tn.Right,
				level: nextLevel,
			})
		}

		// remove current node from queue
		queue.Remove(node)
	}

	// iterate tmp and calculate in res avg value for any level
	res := []float64{tmp[0].sum}
	for i := 1; i < level; i++ {
		res = append(res, tmp[i].sum/tmp[i].count)
	}

	return res
}
