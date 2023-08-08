// package main
package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

func ConcatParams(args []string) string {
	var nstring string
	for index, word := range args {
		nstring += string(word)

		if index != len(args)-1 {
			nstring += "\n"
		}

	}
	return nstring
}

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(ConcatParams(test))
// }
