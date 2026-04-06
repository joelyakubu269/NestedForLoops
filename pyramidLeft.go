package main

import "fmt"

func main() {
	pyramids(5)
}
func pyramids(rows int) {
	for i := 1; i <= rows; i++ {
		for k := 0; k < rows-i; k++ { // this extra loop is used to control when you print from a position other than the left
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
