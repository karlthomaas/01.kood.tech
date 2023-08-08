// package main

package piscine

// import (
// 	"fmt"
// )

func SplitWhiteSpaces(s string) []string {
	var lst []string
	var word string

	for index, char := range s {

		// Characters
		if char > 32 {
			word += string(char)
		}

		// Whitesoaces
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

// func main() {
// 	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
// }
