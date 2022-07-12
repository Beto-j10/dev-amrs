package ex1

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

// Ex1 prints the array of strings in alphabetical order and the number of names
func Ex1() {

	fmt.Println("Ejercicio 1")
	fmt.Println("Ingrese una lista de nombres separados por comas:")

	var strNames string
	fmt.Scanln(&strNames)

	namesArray, namesQuantity := names(strNames)

	fmt.Printf("\nNombres ordenados: %s\n", namesArray)
	fmt.Printf("\nCantidad de nombres: %d\n\n", namesQuantity)
}

// names returns the array of strings and the number of names
func names(strNames string) ([]string, int) {

	namesArray := strings.Split(strNames, ",")
	namesQuantity := len(namesArray)
	sortNames(namesArray)

	return namesArray, namesQuantity

}

// sortNames sorts the array of strings in alphabetical
func sortNames[T constraints.Ordered](x []T) {
	n := len(x)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
