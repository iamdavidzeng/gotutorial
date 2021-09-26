package algorithm

/*
选择排序

分为两层循环，
第一层循环控制选择的次数，
第二层循环控制选出最小值的索引，
选到最小值后和当前的索引交换值。
*/
func SelectionSort(lst []int) []int {
	length := len(lst)
	for i := 0; i < length; i++ {
		pivot := i
		for j := i; j < length; j++ {
			if lst[pivot] > lst[j] {
				pivot = j
			}
		}
		if pivot != i {
			lst[i], lst[pivot] = lst[pivot], lst[i]
		}
	}
	return lst
}
