package main

import "fmt"

func main() {
	sl1 := []int{1, 4, 5, 6, 45, 98, 34, 72}
	bubbleSort(sl1)
}
func bubbleSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(sli)-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j] // bubble sort method compares and pushes the largest number to the end of the slice
			}
		}
	}
	fmt.Println(sli)
}
