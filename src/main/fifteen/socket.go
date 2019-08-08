package main

import (
	"fmt"
	"net"
	"io"
)

func main() {
	var (
		host = "www.apache.org"
		port = "80"
		remote = host + ":" + port
		msg string = "GET / \n"
		data = make([]uint8, 4096)
		read = true
	)

	conn, err := net.Dial("tcp", remote)
	if err != nil {
		return
	}
	io.WriteString(conn, msg)

	for read {
		count, err := conn.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	conn.Close()
}
