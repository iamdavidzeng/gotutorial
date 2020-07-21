package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	timestamp := time.Now().Unix()

	timestampString := strconv.FormatInt(timestamp, 10)
	fmt.Println(reflect.TypeOf(timestamp))

	fmt.Println(timestampString)

	fmt.Println(reflect.TypeOf(timestampString))

	then := time.Date(
		2019, 9, 25, 23, 28, 00, 123456789, time.UTC,
	)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
