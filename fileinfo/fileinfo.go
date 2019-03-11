package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileMode := file.Mode()
	fileSize := file.Size()
	fileSizeB := fileSize / 8
	fileSizeKiB := fileSizeB / 1024
	fileSizeMiB := fileSizeKiB / 1024
	fileSizeGiB := fileSizeMiB / 1024
	isDir := fileMode.IsDir()
	isReg := fileMode.IsRegular()
	permissionBits := fileMode.Perm()
	isAppendOnly := fileMode&os.ModeAppend != 0
	isDevice := false
	isUnixChar := false
	isUnixBloc := false
	if fileMode&os.ModeDevice != 0 {
		isDevice = true
		if fileMode&os.ModeCharDevice != 0 {
			isUnixChar = true
		} else {
			isUnixBloc = true
		}
	}
	isSymbolicLink := fileMode&os.ModeSymlink != 0

	fmt.Printf("filesize: %d bytes, %d KiB, %d MiB, %d GiB\n", fileSizeB, fileSizeKiB, fileSizeMiB, fileSizeGiB)
	fmt.Println("IS Dir ", isDir)
	fmt.Println("Is Reg ", isReg)
	fmt.Println("Uinx code ", permissionBits)
	fmt.Println("is append only: ", isAppendOnly)
	fmt.Println("is device file: ", isDevice)
	fmt.Println("is Unix character device: ", isUnixChar)
	fmt.Println("is Unix block device", isUnixBloc)
	fmt.Println("is symbolic link", isSymbolicLink)
}
