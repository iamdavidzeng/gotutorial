package main

import "fmt"

var res = [][]int{}

func permute(nums []int) [][]int {
	track := []int{}
	helper(nums, track)
	return res
}

func helper(nums []int, track []int) {
	if len(nums) == len(track) {
		res = append(res, track)
		return
	}

	for i := 0; i < len(nums); i++ {
		for _, value := range track {
			if nums[i] == value {
				continue
			}
		}

		track = append(track, nums[i])

		helper(nums, track)

		track = track[:len(track)-1]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
