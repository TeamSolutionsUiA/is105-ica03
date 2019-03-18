package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:5000")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)
}
