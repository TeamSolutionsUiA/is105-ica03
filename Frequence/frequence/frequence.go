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
