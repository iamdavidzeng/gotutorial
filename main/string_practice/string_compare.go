package main

import "strings"

func stringCompare() {

	str1 := "100"
	str2 := "100.05"

	common := strings.Compare(str1, str2)
	if common == 0 {
		println("equals")
	} else {
		println("not equals" + string(common))
	}

	println(common)
}


func stringContain() {

	initial_string := "hello, world"
	
	isContain := strings.Contains(initial_string, "hello")

	println(isContain)
}


func stringFindIndex() {
	initial_string := "hello, world again"

	newIndex := strings.Index(initial_string, ",")
	println(newIndex)
}


func stringSlice() {

	initial_string := " hello , here contanis   many blank.  "

	noBlankString := strings.Fields(initial_string)

	println(noBlankString)

	waitForSliceString := "q,w,e,r,t,y"
	
	slice1 := strings.Split(waitForSliceString, ",")
	println(slice1)

	for index, value := range slice1 {
		println(index, value)
	}

	joinedString := strings.Join(slice1, ".")
	println(joinedString)
}


func main() {
	stringSlice()
}