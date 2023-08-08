package piscine

func RecursivePower(nb int, power int) int {
	if nb > 20 || nb < -20 {
		return 0
	}

	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	if power == 1 {
		return nb
	}

	return nb * RecursivePower(nb, power-1)
}

// func main() {
// 	fmt.Println(RecursivePower(-5, 0))
// }
