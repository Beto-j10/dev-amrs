package main

import (
	"dev/ex1"
	"dev/ex2"
	"dev/ex3"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	for {

		menu := `Seleccione una opción:

	1. ejercicio 1
	2. ejercicio 2
	3. ejercicio 3
	0. salir
`
		fmt.Println(menu)

		var option int
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			ex1.Ex1()
		case 2:
			err := ex2.Ex2()
			if err != nil {
				return err
			}
		case 3:
			ex3.Ex3()
		case 0:
			return nil
		default:
			fmt.Println("Opción inválida")
		}

	}

}
