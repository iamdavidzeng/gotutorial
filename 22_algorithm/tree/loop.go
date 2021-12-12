package main

import (
	"fmt"
)

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func traverse(root *TreeNode) {
	if root.value == 0 {
		return
	}

	fmt.Println(root.value)
	if root.left != nil {
		traverse(root.left)
	}
	if root.right != nil {
		traverse(root.right)
	}
}

func oneSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := oneSideMax(root.left)
	rightMax := oneSideMax(root.right)
	if leftMax > rightMax {
		return leftMax + root.value
	}
	return rightMax + root.value
}

func main() {
	root := &TreeNode{
		value: 5,
		left: &TreeNode{
			value: 1,
		},
		right: &TreeNode{
			value: 2,
			left: &TreeNode{
				value: 3,
			},
			right: &TreeNode{
				value: 4,
			},
		},
	}

	traverse(root)

	fmt.Println(oneSideMax(root))
}
