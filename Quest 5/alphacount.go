// package main
package piscine

// import (
// 	"fmt"
// )

func AlphaCount(s string) int {
	// alphabet lower characters range from 97 - 122
	// alphabet upper characters range from 65 - 90
	charbytes := []byte(s)
	count := 0

	for _, char := range charbytes {
		if (char >= 97 && char <= 122) || (char >= 65 && char <= 90) {
			count++
		}
	}

	return count
}

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := AlphaCount(s)
// 	fmt.Println(nb)
// }
