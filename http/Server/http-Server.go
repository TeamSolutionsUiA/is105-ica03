package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func info() []byte {
	info := Info{"Are", "arefjerm@gmail.com"}
	infoJSON, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	return infoJSON
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(info())
	})

	http.ListenAndServe(":3001", nil)
}

type Info struct {
	Navn  string
	Epost string
}
