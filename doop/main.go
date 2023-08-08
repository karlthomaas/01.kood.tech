package main

import (
	"os"
)

func Atoi(s string) int {
	srune := []rune(s)
	result := 0
	// integer unicode stays between 48-57
	// - 45
	// + 43
	negative_state := false
	for i := range srune {
		if i == 0 {
			if int(srune[i]) == 45 {
				negative_state = true
			} else if int(srune[i]) == 43 {
			} else {
				result = result*10 + int(srune[i]) - '0'
			}
		} else if int(srune[i]) <= 57 && int(srune[i]) >= 48 {
			result = result*10 + int(srune[i]) - '0'
		} else {
			return 696969
		}
	}
	if negative_state == true {
		return -result
	} else {
		return result
	}
}

var runeResult []rune

func createResults(n int) {
	if n < 0 {
		runeResult = append(runeResult, '-')
		n = n * -1
	}
	number := '0'
	if n == '0' {
		runeResult = append(runeResult, '0')
		return
	}
	for var1 := 1; var1 <= n%10; var1++ {
		number++
	}
	for var1 := -1; var1 >= n%10; var1-- {
		number++
	}
	if n/10 != 0 {
		createResults(n / 10)
	}
	runeResult = append(runeResult, number)
}

func main() {
	if len(os.Args) == 4 {
		// Convert's string numbers into real numbers
		a := Atoi(os.Args[1])
		if a == 696969 {
			return
		}
		b := Atoi(os.Args[3])
		if b == 696969 {
			return
		}
		operator := os.Args[2]
		switch operator {
		case "+":
			result := a + b
			if (result > a) == (b > 0) {
				createResults(result)
				runeResult = append(runeResult, 10)
				os.Stdout.Write([]byte(string(runeResult)))

				return
			}
		case "-":
			result := a - b
			if (result < a) == (b > 0) {
				createResults(result)
				runeResult = append(runeResult, 10)
				os.Stdout.Write([]byte(string(runeResult)))
			}
		case "*":
			result := a * b
			if a == 0 || (result/a == b) {
				createResults(result)
				runeResult = append(runeResult, 10)
				os.Stdout.Write([]byte(string(runeResult)))
			}
		case "/":
			if b == 0 {
				os.Stdout.Write([]byte("No division by 0\n"))
			} else {
				result := a / b
				createResults(result)
				runeResult = append(runeResult, 10)
				os.Stdout.Write([]byte(string(runeResult)))
			}
		case "%":
			if b == 0 {
				os.Stdout.Write([]byte("No modulo by 0\n"))
			} else {
				result := a % b
				createResults(result)
				runeResult = append(runeResult, 10)
				os.Stdout.Write([]byte(string(runeResult)))
			}
		}
	}
}
