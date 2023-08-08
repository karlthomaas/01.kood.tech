package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printContent(msg string, state int) {
	for _, char := range msg {
		z01.PrintRune(char)
	}
	// if state == 1 {
	// 	z01.PrintRune('\n')
	// }
}

func main() {
	if len(os.Args) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	} else {
		for _, arg := range os.Args[1:] {
			b, err := ioutil.ReadFile(arg)
			if err != nil {
				printContent("ERROR: "+err.Error(), 0)
				z01.PrintRune('\n')
				os.Exit(1)
			}
			printContent(string(b), 1)
		}
	}
}
