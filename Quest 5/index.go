// package main
package piscine

func Index(s string, toFind string) int {
	lock_state := false
	first_index := -1
	last_index := -1
	match_count := 0

	if toFind == "" {
		return 0
	}
	for _, char := range toFind {
		for index2, char2 := range s {
			// Prevents searching characters before first match
			if lock_state == true {
				if index2 <= last_index {
					continue
				}
			}

			// if there's a character match
			if char == char2 {

				// if it's first match ->
				if lock_state == false {

					// returns if there is only one letter to match
					if len(toFind) == 1 {
						return index2
					}

					last_index = index2
					first_index = index2

					lock_state = true // enables the lock_state
					match_count++
					break
				}

				// If its != first match
				if lock_state == true {
					// check's if the matches are in row
					if last_index+1 == index2 {
						match_count++
						last_index = index2

						// checks if there are more characters to be found
						if match_count == len(toFind) {
							return first_index
						}
						// fmt.Println(index2)
						break

					} else {
						return -1
					}
				}

			}

		}
	}
	return -1
}

// func main() {
// 	fmt.Println(Index("$qgbO%WBURgCi!", "$qgbO%WBUR"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }
