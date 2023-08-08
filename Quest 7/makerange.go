// package main
package piscine

// import (
// 	"fmt"
// )

func MakeRange(min, max int) []int {
	if min == max {
		return nil
	}

	if min > max {
		return nil
	}

	mrange := max - min // numbers between max and min
	empty_list := make([]int, mrange)

	zcounter := 0
	for i := min; i < max; i++ {
		empty_list[zcounter] = i
		zcounter++

	}

	return empty_list
}

// func main() {
// 	fmt.Println(MakeRange(5, 10))
// 	fmt.Println(MakeRange(10, 5))
// }
