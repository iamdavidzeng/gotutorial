package main


func isUgly(num int) bool {

	if num <= 0 {
		return false
	}

	newList := []int{2, 3, 5}

	for _, value := range newList{

		for num % value == 0 {
			num /= value
		}
	}

	return num == 1
}

func main() {
	println(isUgly(2))
}