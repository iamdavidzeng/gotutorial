package algorithm

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	} else {
		less := []int{}
		greater := []int{}
		pivot := nums[0]

		for _, value := range nums[1:] {
			if value <= pivot {
				less = append(less, value)
			} else {
				greater = append(greater, value)
			}
		}
		return append(append(QuickSort(less), pivot), QuickSort(greater)...)
	}
}
