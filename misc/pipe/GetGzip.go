package pipe

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
)

func GetGzipped(expression string) string {

	returnString := GetBase64(expression)
	// https://stackoverflow.com/questions/19197874/how-can-i-use-gzip-on-a-string-in-golang
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte("returnString")); err != nil {
		log.Fatal(err)
	}
	if err := gz.Flush(); err != nil {
		log.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		log.Fatal(err)

	}

	returnString = fmt.Sprintf("%X", gz.Header)
	return returnString

}
