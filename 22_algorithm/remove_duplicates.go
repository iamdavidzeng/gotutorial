package algorithm

func removeDuplicates(nums []int) int {

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		}
	}

	return len(nums)
}
