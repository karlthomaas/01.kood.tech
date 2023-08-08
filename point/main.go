package main

import "github.com/01-edu/z01"

// Defines structure
type point struct {
	x int
	y int
}

func setPoint(pointer *point) {
	// takes in the defined object
	// and changes the object values
	pointer.x = 42
	pointer.y = 21
}

func convertRune(integer int) {
	nrune := '0'

	for i := 1; i <= integer%10; i++ {
		nrune++
	}

	if integer/10 != 0 {
		convertRune(integer / 10)
	}

	z01.PrintRune(nrune)
}

func printResult(x, y int) {
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	convertRune(x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	convertRune(y)
	z01.PrintRune('\n')
}

func main() {
	points := point{} // declares "points" object
	setPoint(&points)
	printResult(points.x, points.y)
}
