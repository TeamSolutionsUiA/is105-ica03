package frequence

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func linjeTeller() {

	// Åpner og leser fra fil.
	linjeTeller := 0

	fil, err := os.Open("../../files/pg100.txt")
	if err != nil {
		log.Fatal(err)
	}
	//buffer := bufio.NewReader(fil)
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		linjeTeller++

	}
	fmt.Println("Filen inneholder", linjeTeller, "linjer")
}

func runeTeller() {
	// Hente inn tekstfilen som runer i en array-buffer.
	fil, err := os.Open("../../files/pg100.txt")
	if err != nil {
		log.Fatal(err)
	}
	textArray := bufio.NewScanner(fil)
	textArray.Split(bufio.ScanRunes)
	var alfaArray [0]string
	var tellerArray [0]int
	teller := 0

	for n := 0; len(textArray) < n; n++ {
		for b := 0; len(alfaArray) < b; b++ {
			if textArray[n] == alfaArray[b] {
				tellerArray[b]++
			} else {
				alfaArray[teller] = textArray[n]
				tellerArray[teller]++
				teller++
			}
		}
	}
	return alfaArray
}
