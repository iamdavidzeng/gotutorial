package main

import "fmt"

func main() {

	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]

	if isPresent {
		fmt.Printf("Beijing is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Paris in map? %t\n", isPresent)
	fmt.Printf("Value is %d\n", value)

	// delete an item:
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("Washington's value is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Washington")
	}
}
