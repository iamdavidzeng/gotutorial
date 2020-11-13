package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}

func main() {
	nums := []int{1}

	rotate(nums, 2)

	fmt.Println(nums)
}
