package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		original_nb := nb
		for i := 1; i < power; i++ {
			nb *= original_nb
		}
		return nb
	}
}

// func main() {
// 	fmt.Println(IterativePower(4, 3))
// }
