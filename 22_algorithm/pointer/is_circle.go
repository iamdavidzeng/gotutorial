package main

import "fmt"

type LinkNode struct {
	next *LinkNode
}

func isCircle(head *LinkNode) bool {
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			return true
		}
	}

	return false
}

func main() {
	head := &LinkNode{}
	mid := &LinkNode{}

	head.next = mid
	mid.next = head

	fmt.Println(isCircle(head))
}
