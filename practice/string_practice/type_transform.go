package main

import "strconv"


func stringTransform() {

	initial_int := "100"

	transformed_int, err := strconv.Atoi(initial_int)
	if err != nil {
		println(err)
	} else {
		println(transformed_int)
	}

	initial_float_str := "100.05"

	transformed_float, err := strconv.ParseFloat(initial_float_str, 10)
	if err != nil {
		println(err)
	} else {
		println(transformed_float)
	}

}


func main() {
	stringTransform()
}