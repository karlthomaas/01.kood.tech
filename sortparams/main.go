package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	params_length := len(params)

	var sorted_list []rune

	for {

		if params_length == len(sorted_list) {
			break
		}

		lowest_value := params[0]
		lowest_index := 0

		for index, elem := range params {
			if elem < lowest_value {
				lowest_value = elem
				lowest_index = index
			}
		}

		sorted_list = append(sorted_list, rune(lowest_value[0]))
		params = append(params[:lowest_index], params[lowest_index+1:]...)

	}

	for _, item := range sorted_list {
		z01.PrintRune(item)
		z01.PrintRune('\n')
	}
}
