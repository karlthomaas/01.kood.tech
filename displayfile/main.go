package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return

	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	file_name := args[0]

	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error...")
		return
	}

	arr := make([]byte, 15)
	file.Read(arr)
	fmt.Print(string(arr))
}
