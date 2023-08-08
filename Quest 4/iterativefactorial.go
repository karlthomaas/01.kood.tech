package piscine

func IterativeFactorial(nb int) int {
	if nb < -20 || nb > 20 {
		return 0
	}
	number := 1
	if nb == 0 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			number = number * i
		}
	}
	return number
}
