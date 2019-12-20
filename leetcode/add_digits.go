package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	} else {
		mid := 0
		for num != 0 {
			mid += num % 10
			num /= 10
		}
		return addDigits(mid)
	}
}


func main() {
	res:= addDigits(38)
	fmt.Println(res)
}
