package bruteforce

import (
	"avala/common"
	"fmt"
)

func InitialiseDb(numberOfHexDigits int8) {
	//upperLimit := int(math.Pow(0x10, float64(numberOfHexDigits))) - 1
	db := common.GetDb()
	fmt.Printf("%v\n", db)
}
