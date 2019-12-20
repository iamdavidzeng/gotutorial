package main

var c string

func main() {
	c = "G"
	print(c)
	f1()
}

func f1() {
	c := "0"
	print(c)
	f2()
}

func f2() {
	print(c)
}
