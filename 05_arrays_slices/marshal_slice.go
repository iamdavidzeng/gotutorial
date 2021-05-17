package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	foo := []int{}
	bar := make([]int, 0)
	var baz []int

	output1, _ := json.Marshal(foo)
	output2, _ := json.Marshal(bar)
	output3, _ := json.Marshal(baz)

	fmt.Println(string(output1))
	fmt.Println(string(output2))
	fmt.Println(string(output3))
}
