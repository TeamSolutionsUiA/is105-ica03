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
	fmt.Println("Ukomprimert test:")
	fmt.Println("Content (first 10 bytes): " + hexSlice[0:20] + "\n" +
		"Length: " + hexLength)

	// Test getbase64
	base64String := pipe.GetBase64(fileString)
	base64Slice := base64String[:100]
	base64Length := fmt.Sprintf("%d", len(base64String))
	fmt.Println("base64 test:")
	fmt.Println("Content (first 10 bytes): " + base64Slice[0:20] + "\n" +
		"Length: " + base64Length)

	// Test av GetGzip
	gzipString := pipe.GetGzipped(fileString)
	gzipSlice := gzipString[:100]
	gzipLength := fmt.Sprintf("%d", len(gzipString))
	fmt.Println("Gzip test:")
	fmt.Println("Content (first 10 bytes): " + gzipSlice[0:20] + "\n" +
		"Length: " + gzipLength)
}
