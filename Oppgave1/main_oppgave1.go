package main

import (
	"fmt"

	"github.com/TeamSolutionsUiA/ica02/is105-ica02-master/fileutils/fileutils"
)

func main() {
	text1 := fileutils.FileToByteslice("../files/text1.txt")
	text2 := fileutils.FileToByteslice("../files/text2.txt")
	fmt.Printf("%c \n%c", text1, text2)
	fmt.Printf("%x \n%x", text1, text2)
}
