package main

import "fmt"

func main() {
	pyramid(5)
}
func pyramid(rows int) {
	for i := 1; i <= rows; i++ {
		for k := 0; k < rows-i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
