// package main

// import "fmt"

package piscine

func AppendRange(min, max int) []int {
	empty_list := []int{}

	if min > max {
		return nil
	}

	if min == max {
		return nil
	}

	for i := min; i < max; i++ {
		empty_list = append(empty_list, i)
	}

	return empty_list
}

// func main() {
// 	fmt.Println(AppendRange(5, 10))
// 	fmt.Println(AppendRange(10, 5))
// 	fmt.Println(AppendRange(1, 1))

// }
