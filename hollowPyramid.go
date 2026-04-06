//	   *
//	  * *
//	 *   *
//	*     *
//
// *********
package main

import "fmt"

func main() {
	holPyramid(5)
}
func holPyramid(rows int) {
	for i := 1; i <= rows; i++ {
		for k := 0; k < rows-i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			if j == 0 || j == 2*i-2 || i == rows { // here we use 2*i-2 as the boundary because we start from zero && aim
				// for the last index hence 2*i -1 -1
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
