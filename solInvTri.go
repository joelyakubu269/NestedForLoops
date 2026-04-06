package main

import "fmt"

func main() {
	solInvTri(5)
}
func solInvTri(rows int) {
	for i := 1; i <= rows; i++ {
		for j := rows; j >= i; j-- {
			if i == 1 || i == rows {
				fmt.Print("*")
			} else if j == rows || j == i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
