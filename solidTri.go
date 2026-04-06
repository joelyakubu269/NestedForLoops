package main

import "fmt"

func main() {
	solTri(5)
}
func solTri(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
