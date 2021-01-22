package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		for {
			c := <-ch
			fmt.Println("first", c, "len", len(ch))
		}
	}()

	go func() {
		for {
			c := <-ch
			fmt.Println("first", c, "len", len(ch))
		}
	}()

	for i := 1; i <= 100; i++ {
		ch <- i
	}

	time.Sleep(time.Second * 1)

}
