package picsine

// import (
// 	"github.com/01-edu/z01"
// )

func FirstRune(s string) rune {
	bytes := []rune(s)    // Convert string to runes
	return rune(bytes[0]) // Print first rune
}

// func main() {
// 	z01.PrintRune(FirstRune("Hello!"))
// 	z01.PrintRune(FirstRune("Salut!"))
// 	z01.PrintRune(FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }
