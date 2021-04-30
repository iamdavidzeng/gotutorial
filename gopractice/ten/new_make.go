package main


type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}


func main() {
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	//z := make(Bar)
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	//u := new(Foo)
	//(*u)["x"] = "goodbye"
	//(*u)["y"] = "world"
}