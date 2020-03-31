package main

import "fmt"

func removeDuplicates(nums []int) int {

	for i := 1; i < len(nums); i++ {

		for j := 0; j < i; j++ {

			if nums[i] == nums[j] {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		}
	}

	fmt.Println(nums)

	return len(nums)
}

func main() {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	removeDuplicates(nums)

}
