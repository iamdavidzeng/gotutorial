package main

import (
	"errors"
	"fmt"
	"time"
)

func doConcurrently(job string, err chan error) {
	go func() {
		fmt.Println("doing job", job)
		time.Sleep(1 * time.Second)
		err <- errors.New("somthing went wrong!")
	}()
}

func do(job string) error {
	fmt.Println("doing job", job)
	time.Sleep(1 * time.Second)
	return errors.New("something went wrong!")
}

func main() {
	jobs := []string{"one", "two", "three"}

	errc := make(chan error)
	for _, job := range jobs {
		// doConcurrently(job, errc)
		go func(job string) {
			errc <- do(job)
		}(job)
	}
	for _ = range jobs {
		if err := <-errc; err != nil {
			fmt.Println(err)
		}
	}
}
