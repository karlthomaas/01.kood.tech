// package main

package piscine

// import "fmt"

// import (
// 	"fmt"
// )

func Split(s, sep string) []string {
	sep_state := false
	var indexes []int
	var total_indexes []int
	match_index := -1

	// Starts to iterate through given string
	// and collect's every index where sep has presence

	for index, char := range s {
		// When the character matches with sep's first character
		if sep_state == false && string(char) == string(sep[0]) {
			sep_state = true
			match_index++ // Set's match index as 0
			indexes = append(indexes, index)
			continue
		}

		if sep_state == true {
			if match_index > len(sep) {
				continue
			}

			// if the next chatacter matches with next sep character
			if string(char) == string(sep[match_index+1]) {
				match_index++
				// adds the current index to list (saves all the sep occurences)
				indexes = append(indexes, index)

				// Check's if all of the indexes have been found
				if match_index == len(sep)-1 {
					for _, index := range indexes {
						total_indexes = append(total_indexes, index)
					}
					// Resets stats
					match_index = -1
					sep_state = false
					indexes = []int{}
					sep_state = false
					continue
				}

			} else {
				// if it doesn't match -> resets the current stats
				match_index = -1
				sep_state = false
				indexes = []int{}
			}
		}
	}

	str_lst := []string{}

	presence := false
	var word string
	match_count2 := 0
	for index, char := range s {

		if match_count2 <= len(total_indexes) {
			for _, sindex := range total_indexes {
				if sindex == index {
					match_count2++
					presence = true
				}
			}
		}

		if presence == false {
			word += string(char)
		} else {
			if word != "" {
				str_lst = append(str_lst, word)
			}
			presence = false
			word = ""
		}

	}
	str_lst = append(str_lst, word)
	return str_lst
}

// func main() {
// 	s := "ghSNfBgXqCGMd|=choumi=|2dl7Sjnxb3RRI|=choumi=|rI6cc3r9SYJd8|=choumi=|Rnq0uHAcOf6Bf|=choumi=|WoxoGHDwEgIKH|=choumi=|GK2s4uQKW5q8k|=choumi=|6m56HBdYK12U0|=choumi=|bky55w3wZmj55"
// 	fmt.Printf("%#v\n", Split(s, "|=choumi=|"))
// 	// a := "HelloHAhowHAareHAyou?"
// 	// fmt.Printf("%#v\n", Split(a, "HA"))
// }
