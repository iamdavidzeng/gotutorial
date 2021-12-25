package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 前序遍历
	fmt.Println("link table before:", head.Val)

	traverseList(head.Next)

	// 后续遍历
	fmt.Println("link table after:", head.Val)

	return nil
}

func traverseTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 前序
	traverseTree(root.Left)
	// 中序
	traverseTree(root.Right)
	// 后序
	fmt.Println("tree after:", root.Val)

	return nil
}

func main() {
	// string
	str := "david"
	for i, v := range str {
		fmt.Println(i, string(v))
	}

	// array
	ids := []int{11, 22, 33, 44}
	for i, v := range ids {
		fmt.Println(i, v)
	}

	// slice
	newIds := ids[:len(ids)-1]
	for i, v := range newIds {
		fmt.Println(i, v)
	}

	// map
	emails := map[string]string{
		"David": "david@gmail.com",
		"Bob":   "bob@gmail.com",
	}

	for k, v := range emails {
		fmt.Println(k, v)
	}

	// link table iteration
	linkTable := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	for p := linkTable; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

	// link table recursion
	traverseList(linkTable)

	// binary tree recursion
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	traverseTree(tree)
}
