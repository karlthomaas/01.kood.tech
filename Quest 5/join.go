package piscine

// package main

// import (
// 	"fmt"
// )

func Join(strs []string, sep string) string {
	nstring := ""
	for index, word := range strs {
		nstring += string(word)
		if len(strs)-1 != index {
			nstring += sep
		}
	}
	return nstring
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(Join(toConcat, ":"))
// }
