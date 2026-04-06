package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str1 := "A man, a plan, a canal: Panama"
	fmt.Println(PalUpgrade(str1))

}
func PalUpgrade(str string) bool {
	word := strings.ToLower(str)
	cleaned := []rune{}
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}
	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}
	return true
}
