package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ipAdresser := []string{"10.228.35.33:5000", "10.228.3.106:5000", "10.228.3.122:5000", "10.228.1.176:5000", "10.228.19.77:5000"}
	for _, ip := range ipAdresser {
		conn, _ := net.Dial("tcp", ip)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message + "\n")
	}
}
