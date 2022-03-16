package main

import "fmt"

func removeDuplicates(nums []int) int {
	p1 := 1

	for p2 := 1; p2 < len(nums); p2++ {
		if nums[p2] != nums[p2-1] {
			nums[p1] = nums[p2]
			p1++
		}
		fmt.Println(p1, p2, nums)
	}
	return p1
}

func main() {
	removeDuplicates([]int{0, 1, 1, 2})
}
