package frequence

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func linjeTeller() {

	
	linjeTeller := 0

	// Åpner filen som vi ønsker å telle linjer i.
	fil, err := os.Open("../../files/pg100.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Lager en scanner som kan finne linjeskift i teksten og telle dem.
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		linjeTeller++

	}
	// Printe ut resultatet.
	fmt.Println("Filen inneholder", linjeTeller, "linjer")
}

// Trenger en funksjon for å telle antall ord (eller runes) i tekst-dokumentet.

func runeTeller() {
	
}
