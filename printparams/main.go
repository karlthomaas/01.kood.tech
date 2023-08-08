package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, a := range os.Args[1:] {
		for _, c := range a {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
