package main

import (
	"fmt"
	"sort"
)

type Node struct {
	index int
	value int
}

func advantageCount(nums1 []int, nums2 []int) []int {
	nodes := []Node{}
	for i, v := range nums2 {
		nodes = append(nodes, Node{index: i, value: v})
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].value < nodes[j].value
	})
	sort.Ints(nums1)

	fmt.Println(nodes)

	res := make([]int, len(nums1))
	left, right := 0, len(nums2)-1
	for i := len(nums2) - 1; i >= 0; i-- {
		if nums1[right] > nodes[i].value {
			res[nodes[i].index] = nums1[right]
			right -= 1
		} else {
			res[nodes[i].index] = nums1[left]
			left += 1
		}
	}

	return res
}

func main() {
	nums1 := []int{8, 2, 4, 4, 5, 6, 6, 0, 4, 7}
	nums2 := []int{0, 8, 7, 4, 4, 2, 8, 5, 2, 0}
	fmt.Println(advantageCount(nums1, nums2))
}
