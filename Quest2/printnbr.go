package piscine

import "github.com/01-edu/z01"

func Run(n int) {
	number := '0'
	if n == '0' {
		z01.PrintRune('0')
		return
	}
	for var1 := 1; var1 <= n%10; var1++ {
		number++
	}
	for var1 := -1; var1 >= n%10; var1-- {
		number++
	}
	if n/10 != 0 {
		Run(n / 10)
	}
	z01.PrintRune(number)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	Run(n)
}
