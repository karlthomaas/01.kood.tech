package piscine

// package main

// import (
// 	"fmt"
// )

func BasicJoin(elems []string) string {
	nstring := ""
	for _, word := range elems {
		nstring += string(word)
	}

	return nstring
}

// func main() {
// 	elems := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(BasicJoin(elems))
// }
