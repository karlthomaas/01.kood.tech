package piscine

// package main

func IsSorted(f func(a, b int) int, a []int) bool {
	state := true
	increasing := true

	if len(a) == 1 || len(a) == 0 {
		return state
	}

	if a[0] > a[1] {
		increasing = false
	}

	for index, nr := range a {

		index += 1
		if index > len(a)-1 {
			return state
		}

		condition := f(a[index], nr)

		if condition == 0 {
			continue
		}

		if increasing == true && !(condition > 0) || increasing == false && !(condition < 0) {
			state = false
			return state
		}
	}
	return state
}

// func f(a, b int) int {
// 	if a > b {
// 		return 1
// 	} else {
// 		return -1
// 	}
// }
// func main() {
// 	a1 := []int{0, -2, -5, -10, -15}
// 	a2 := []int{0, 2, 1, 3}

// 	result1 := IsSorted(f, a1)
// 	result2 := IsSorted(f, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
