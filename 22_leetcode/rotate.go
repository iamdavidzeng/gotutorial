package main

import "fmt"

func rotate(nums []int, k int) []int {
	i := len(nums) - k
	nums = append(nums[i:], nums[:i]...)
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(rotate(nums, 3))
}
