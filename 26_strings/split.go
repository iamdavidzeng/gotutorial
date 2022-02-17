package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "health_check"
	slc := strings.Split(input, "_")
	for i := range slc {
		slc[i] = strings.Title(slc[i])
	}
	val := strings.Join(slc, "")
	fmt.Println(val)

	for i, v := range input {
		fmt.Println(i, string(v))
	}
}
