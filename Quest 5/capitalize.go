package piscine

// package main

// import (
// 	"fmt"
// )

func checkAlphaNumeric(char byte) bool {
	// if string(char) < "a" || string(char) > "z" {

	if !(string(char) >= "a" && string(char) <= "z") && !(string(char) >= "0" && string(char) <= "9") && !(string(char) >= "A" && string(char) <= "Z") {
		return true
	} else {
		return false
	}
}

func Capitalize(s string) string {
	bytes := []byte(s)
	nstring := ""

	for index, char := range bytes {
		if index == 0 && (string(bytes[0]) >= "a" && string(bytes[0]) <= "z") {
			bytes[0] = bytes[0] - 32
		}

		if string(char) >= "A" && string(char) <= "Z" {
			if index != 0 {
				bytes[index] = char + 32
			}
		}
	}

	for index, char := range bytes {
		if checkAlphaNumeric(char) {
			if index == len(s)-1 {
				nstring += string(char)
				return nstring
			}
			nchar := bytes[index+1]
			if string(nchar) >= "a" && string(nchar) <= "z" {
				bytes[index+1] = nchar - 32
			}
		}

		nstring += string(char)
	}

	return nstring
}

// // Capitalize("{#o@;.o!c.uip") == "{#o@;.o!c.uip" instead of "{#O@;.O!C.Uip"
// Capitalize("Mq[,O~u19H:2,") == "Mq[,O~U19h:2" instead of "Mq[,O~U19h:2,"

// func main() {
// 	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
// 	fmt.Println(Capitalize("{#o@;.o!c.uip"))
// 	fmt.Println(Capitalize("Mq[,O~u19H:2,"))
// }
