package main

import (
	"encoding/json"
	"log"
	"net"
)

func handler(c net.Conn) {
	info := Info{"Are", "arefjerm@gmail.com"}
	infoJSON, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	c.Write([]byte(infoJSON))
	c.Close()
}
func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}

type Info struct {
	Navn  string
	Epost string
}
