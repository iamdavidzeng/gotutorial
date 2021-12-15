package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}

	var helper func(nums []int, track []int)
	helper = func(nums []int, track []int) {
		if len(nums) == len(track) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if contains(track, nums[i]) {
				continue
			}

			track = append(track, nums[i])

			helper(nums, track)

			track = track[:len(track)-1]
		}
	}

	track := []int{}
	helper(nums, track)
	return res

}

func contains(lst []int, value int) bool {
	for _, v := range lst {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(permute([]int{5, 4, 6, 2}))
}
