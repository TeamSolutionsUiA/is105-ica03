package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	responce, err := http.Get("http://127.0.0.1:3001/")
	if err != nil {
		log.Fatal(err)
	}
	defer responce.Body.Close()
	body, err := ioutil.ReadAll(responce.Body)
	fmt.Println("get:\n", string(body))
}
