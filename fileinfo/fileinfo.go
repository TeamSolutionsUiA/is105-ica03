package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Lstat("../files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileMode := file.Mode()
	fileSize := file.Size()
	isDir := fileMode.IsDir()
	isReg := fileMode.IsRegular()
	permissionBits := fileMode.Perm()
	//isAppendOnly :=
	//isDeviseFile :=
	//isUnixChar :=
	//isUnixBloc :=
	//isSymbolicLink :=
	fmt.Print(fileSize, isDir, isReg, permissionBits)
}
