// package main
package piscine

// import (
import "github.com/01-edu/z01"

// )

func SplitWhiteSpaces(s string) []string {
	var lst []string
	var word string

	for index, char := range s {

		// Characters
		if char > 32 {
			word += string(char)
		}

		// Whitespaces
		if char <= 32 {
			if len(word) != 0 {
				lst = append(lst, word)
			}
			word = ""
		}

		if index == len(s)-1 {
			lst = append(lst, word)
			return lst
		}

	}
	return lst
}

func PrintWordsTables(a []string) {
	for _, elem := range a {
		for _, char := range elem {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	a := SplitWhiteSpaces("Hello how are you?")
// 	PrintWordsTables(a)
// }
