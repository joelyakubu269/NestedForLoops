package main

import (
	"fmt"
)

func main() {
	holSquare(5, 5)
}
func holSquare(rows, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || i == rows-1 {
				fmt.Print("*")
			} else if j == 0 || j == cols-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
