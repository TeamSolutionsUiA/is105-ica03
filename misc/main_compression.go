package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TeamSolutionsUiA/is105-ica03/misc/pipe"
)

func main() {
	filename := os.Args[1]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileString := string(file)

	// Test getHex
	hexString := pipe.GetHex(fileString)
	fmt.Println("Content: " + hexString + "\n")

	// Test getbase64
	base64String := pipe.GetBase64(fileString)
	fmt.Println("Content: " + base64String + "\n")

	// Test av GetGzip
	gzipString := pipe.GetGzipped(fileString)
	fmt.Println("Header content: " + gzipString)
}
