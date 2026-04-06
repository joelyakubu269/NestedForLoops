package main

import "fmt"

func main() {
	holTri(5)
}
func holTri(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			if i == 1 || i == rows {
				fmt.Print("*")
			} else if j == 1 || j == i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
