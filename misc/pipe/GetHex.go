package pipe

import (
	"fmt"
)

func GetHex(expression string) string {

	returnString := expression
	returnString = fmt.Sprintf("%X", returnString)

	return returnString
}
