package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			x := j + 1
			for k := i; k <= '9'; k++ {
				for ; x <= '9'; x++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(x)
					if i == '9' && j == '8' && k == '9' && x == '9' {
						// do nothing > python: ...
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				x = '0'
			}

		}
	}
	z01.PrintRune('\n')
}
