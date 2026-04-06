package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "madam"
	fmt.Println(Palindrome(str1))
}
func Palindrome(str string) bool {
	word := strings.ToLower(str)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
