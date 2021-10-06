package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type NewFoo struct {
	Name string
}

func (f *NewFoo) Hello() (map[string]string, error) {
	return map[string]string{
		"git_sha": f.Name,
	}, nil
}

func main() {

	f := &NewFoo{
		Name: "foo",
	}

	fn := reflect.ValueOf(f).MethodByName("Hello")
	fmt.Println(fn)
	val := fn.Call([]reflect.Value{})

	fmt.Println(fn.Type())

	result, _ := json.Marshal(val)
	fmt.Println(string(result))

}
