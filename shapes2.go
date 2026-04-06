package main

import (
	"fmt"
)

func main() {
	conCat(25)
}
func conCat(rows int) {
	for i := 1; i <= rows; i++ {
		str := ""

		for j := 1; j <= i; j++ {
			str += "G"

		}
		fmt.Println(str)

	}
}
