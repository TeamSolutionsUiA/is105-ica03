package lineshift

import (
	"fmt"

	"github.com/TeamSolutionsUiA/ica02/is105-ica02-master/fileutils/fileutils"
)

func LineshiftType(filename string) string {
	byteSlice := fileutils.FileToByteslice(filename)
	for i := 0; i < len(byteSlice); i++ {
		currentByte := fmt.Sprintf("%X", byteSlice[i])
		if currentByte == "D" {
			nextByte := fmt.Sprintf("%X", byteSlice[i+1])
			if nextByte == "A" {
				return "Denne filen bruker Windows linjeskjift(CR LF)"
			}
			return "Denne filen bruker linjeskjift for Commondore 8-bit machines (CR)"
		} else if currentByte == "A" {
			nextByte := fmt.Sprintf("%X", byteSlice[i+1])
			if nextByte == "D" {
				return "Denne filen bruker Acorn BBC/Risc OS linjeskjift(LF+CR)"
			}
			return "Denne filen bruker Linux/Mac linjeskjift (LF)"
		}

	}

	return "ERROR"
}
