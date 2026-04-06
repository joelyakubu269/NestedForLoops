package main

import (
	"fmt"
)

func main() {
	Squre(5, 5)
}
func Squre(rows, cols int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
