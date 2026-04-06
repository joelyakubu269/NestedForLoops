package main

import (
	"fmt"
)

func main() {
	conCat(25)
}
func conCat(rows int) {
	str := ""
	for i := 1; i <= rows; i++ {

		str += "G"
		fmt.Println(str)

	}
}
