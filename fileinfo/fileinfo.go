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
	appendhex:="2d72772d72772d72772d";
	fileMode := file.Mode()
	fileSize := file.Size()
	isDir := fileMode.IsDir()
	isReg := fileMode.IsRegular()
	permissionBits := fileMode.Perm()
	isAppendOnly :=false;
	if (fmt.Sprintf("%x",permissionBits)==appendhex){
		isAppendOnly=true;

	}
	
	

	//isDeviseFile :=
	
	//isUnixChar :=
	
	//isUnixBloc :=
	
	//isSymbolicLink :=
	
	fmt.Print("filesize byte ", fileSize, )
	fmt.Print("IS Dir ", isDir)
	fmt.Print("Is Reg ", isReg )
	fmt.Print("Uinx code ", permissionBits) 
	fmt.Print("is append only ", isAppendOnly)


	
	
}
