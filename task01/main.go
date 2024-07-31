package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		strValue = "104"
		intValue = 35
	)
	convertStr, err := strconv.Atoi(strValue)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(convertStr)
	}

	convertInv := strconv.Itoa(intValue)
	fmt.Println(convertInv)

}
