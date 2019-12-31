package main

import "fmt"

func main() {
	// Long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Short method
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("Num %d is FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("Num %d is Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("Num %d is Buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
