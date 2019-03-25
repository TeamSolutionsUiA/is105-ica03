package main

import (
	"fmt"

	"github.com/TeamSolutionsUiA/is105-ica03/Frequence/bfrequence"
	"github.com/TeamSolutionsUiA/is105-ica03/Frequence/frequence"
)

func main() {

	fmt.Println("frequence test: ")

	frequence.LinjeTeller()
	frequence.RuneTeller()

	fmt.Println("bfrequence test: ")

	bfrequence.LinjeTeller()
	bfrequence.RuneTeller()
}
