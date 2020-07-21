package main

import (
	"fmt"
	"strconv"
)

// uint8  : 0 to 255
// uint16 : 0 to 65535
// uint32 : 0 to 4294967295
// uint64 : 0 to 18446744073709551615
// int8   : -128 to 127
// int16  : -32768 to 32767
// int32  : -2147483648 to 2147483647
// int64  : -9223372036854775808 to 9223372036854775807

// int8's bitSize is 8
func intEightMaxValue(s string) {
	i, err := strconv.ParseInt(s, 10, 8)
	fmt.Println(i, err)
}

func main() {
	intEightMaxValue("129")
}
