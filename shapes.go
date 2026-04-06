package main

import (
	"fmt"
)

func main() {
	gRows(25)
}

//	func gRows(rows int) {
//		for i := 1; i <= rows; i++ {
//			for j := 1; j <= i; j++ {
//				fmt.Print("G")
//			}
//			fmt.Println()
//		}
//	}
//
// An alternative method
func gRows(rows int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
}
