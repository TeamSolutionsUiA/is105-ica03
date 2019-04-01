package pipe

import (
	"encoding/base64"
	"fmt"
)

func GetBase64(expression string) string {

	returnString := expression
	encoded := base64.StdEncoding.EncodeToString([]byte(returnString))

	returnString = fmt.Sprintf("%X", encoded)
	return returnString
}
