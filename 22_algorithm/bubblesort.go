package algorithm

func BubbleSort(lst []int) []int {
	length := len(lst)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if lst[j] > lst[j+1] {
				lst[j], lst[j+1] = lst[j+1], lst[j]
			}
		}
	}

	return lst
}
