package main

import "fmt"

func main() {
	sl1 := []int{1, 23, 4, 56, 34, 89, 46, 5, 7}
	selectionSort(sl1)

}
func selectionSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		min := i
		for j := i + 1; j < len(sli); j++ { // for each position i compares other values ahead i+1
			if sli[j] < sli[min] { // in contrast to the bubble sort here we store positions at minimum values
				min = j // store the index
			}
		}
		sli[i], sli[min] = sli[min], sli[i]
	}
	fmt.Println(sli)
}
