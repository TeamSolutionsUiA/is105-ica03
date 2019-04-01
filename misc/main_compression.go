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
	hexSlice := hexString[:100]
	hexLength := fmt.Sprintf("%d", len(hexString))
	fmt.Println("Content: " + hexSlice + "\n" +
		"Length: " + hexLength)

	// Test getbase64
	base64String := pipe.GetBase64(fileString)
	base64Slice := base64String[:100]
	base64Length := fmt.Sprintf("%d", len(base64String))
	fmt.Println("Content: " + base64Slice + "\n" +
		"Length: " + base64Length)

	// Test av GetGzip
	gzipString := pipe.GetGzipped(fileString)
	gzipSlice := gzipString[:100]
	gzipLength := fmt.Sprintf("%d", len(gzipString))
	fmt.Println("Content: " + gzipSlice + "\n" +
		"Length: " + gzipLength )
}
