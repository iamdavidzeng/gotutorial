package algorithm

func InsertionSort(lst []int) []int {
	length := len(lst)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j > 0; j-- {
			if lst[j] < lst[j-1] {
				lst[j], lst[j-1] = lst[j-1], lst[j]
			}
		}
	}

	return lst
}
