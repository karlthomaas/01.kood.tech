package main

import "fmt"

func Sort(lst []int) []int {
	list_length := len(lst)
	var new_list []int
	for {
		if len(new_list) == list_length {
			break
		}
		lowest_value := lst[0]
		lowest_index := 0
		for index, integer := range lst {
			if integer < lowest_value {
				lowest_value = integer
				lowest_index = index
			}
		}
		new_list = append(new_list, lowest_value)
		lst = append(lst[:lowest_index], lst[lowest_index+1:]...)
		fmt.Println(new_list)
		fmt.Println(lst)
	}
	return new_list
}

func main() {
	int_lst := []int{9, 8, 7, 10, 2, 1}
	fmt.Print(Sort(int_lst))
}
