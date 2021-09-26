package algorithm

func IsUgly(num int) bool {

	initList := []int{2, 3, 5}
	for _, value := range initList {
		for num%value == 0 {
			num /= value
		}
	}
	return num == 1
}
