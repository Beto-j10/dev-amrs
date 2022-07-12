package ex3

import (
	"fmt"
	"strings"
)

// Ex3 prints if the two strings are friends
func Ex3() {
	fmt.Println("Ejercicio 3")
	fmt.Println("Ingrese cadena 1:")
	var str1 string
	fmt.Scanln(&str1)
	fmt.Println("Ingrese cadena 2:")
	var str2 string
	fmt.Scanln(&str2)
	fmt.Printf("\nSon amigos: %t\n\n", checkFriends(str1, str2))

}

// checkFriends returns true if the two strings are friends
func checkFriends(str1, str2 string) bool {
	lenStr1 := len(str1)
	lenStr2 := len(str2)

	if lenStr1 != lenStr2 {
		return false
	}

	str1Array := strings.Split(str1, "")

	count := 0

	var rotate func(str []string) bool

	rotate = func(str []string) bool {

		count++
		if count > lenStr1 {
			return false
		}

		if strings.Join(str, "") == str2 {
			return true
		}

		return rotate(append(str, str[0])[1:])
	}

	return rotate(str1Array)
}
