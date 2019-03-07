package main

import (
	"fmt"
	"os"

	"github.com/TeamSolutionsUiA/is105-ica03/lineshift/lineshift"
)

func main() {
	filename := os.Args[1]
	returnString := lineshift.LineshiftType(filename)
	fmt.Print(returnString)
}
